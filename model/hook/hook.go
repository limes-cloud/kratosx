package hook

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	Read   = "read"
	Create = "create"
	Update = "update"
	Delete = "delete"

	TagHook    = "HOOK"
	TypeTenant = "tenant"
	TypeDept   = "dept"
	TypeUser   = "user"
)

type op struct {
	check bool
	set   bool
	where bool
}

type column struct {
	name string
	db   string
	op   map[string]*op
}

type Condition struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// ConditionGroup 定义条件组，可以包含多个条件和子条件组
type ConditionGroup struct {
	Logic      string            `json:"logic"`      // "AND" 或 "OR"
	Conditions []*Condition      `json:"conditions"` // 子条件列表
	Groups     []*ConditionGroup `json:"groups"`     // 子条件组列表
}

type ScopeRequestFunc func(ctx context.Context, database string, model string, method string) (ScopeResponse, error)

type ScopeResponse interface {
	// DeptScopes 部门权限
	DeptScopes() (bool, []uint32)
	// UserId 用户id
	UserId() uint32
	// DeptId 用户部门
	DeptId() uint32
	// TenantId 当前的租户ID
	TenantId() uint32
	// UserDeptId 指定用户的主部门
	UserDeptId(uint32) uint32
	// Condition 条件约束
	Condition() *ConditionGroup
	// Fields 查询字段
	Fields() []string
}

var tagReg = regexp.MustCompile(`\[(.*?)\]`)

func Apply(d *gorm.DB, dbName, method string, req ScopeRequestFunc) {
	if req == nil {
		return
	}

	// 获取权限数据
	resp, err := req(d.Statement.Context, dbName, d.Statement.Table, method)
	if err != nil {
		_ = d.AddError(err)
		return
	}

	if resp == nil {
		return
	}

	hfs := map[string][]column{}
	if d.Statement.Schema != nil {
		for _, field := range d.Statement.Schema.Fields {
			tv, ok := field.TagSettings[TagHook]
			key, opv := parseTag(tv)
			if ok {
				hfs[key] = append(hfs[key], column{
					name: field.Name,
					db:   field.DBName,
					op:   opv,
				})
			}
		}
	}

	// 没有hook字段，则直接跳过
	if len(hfs) == 0 {
		applyJoinQuery(d, resp)
		return
	}

	switch method {
	case Create:
		applyCheckAndSet(d, Create, hfs, resp)
	case Update:
		applyCheckAndSet(d, Update, hfs, resp)
		applyWhere(d, Update, hfs, resp)
		applyCondition(d, resp)
	case Read:
		applySelect(d, resp)
		applyWhere(d, Read, hfs, resp)
		applyCondition(d, resp)
	case Delete:
		applyWhere(d, Delete, hfs, resp)
		applyCondition(d, resp)

	}
}

func applyCheckAndSet(d *gorm.DB, method string, hfs map[string][]column, resp ScopeResponse) {
	for tp, cols := range hfs {
		var err error
		for _, col := range cols {
			err = checkAndSetValue(d.Statement.ReflectValue, method, tp, col, resp)
			if err != nil {
				_ = d.AddError(err)
				return
			}
		}

	}
}

func applyWhere(d *gorm.DB, method string, hfs map[string][]column, resp ScopeResponse) {
	all, scope := resp.DeptScopes()
	for tp, cols := range hfs {
		for _, col := range cols {
			if !col.op[method].where {
				continue
			}
			switch tp {
			case TypeTenant:
				d = d.Where(fmt.Sprintf("`%s`.`%s` = ?", d.Statement.Table, col.db), resp.TenantId())
			case TypeDept:
				if !all && len(scope) > 0 {
					d = d.Where(fmt.Sprintf("`%s`.`%s` in ?", d.Statement.Table, col.db), scope)
				}
			case TypeUser:
				if !all && len(scope) == 0 {
					d = d.Where(fmt.Sprintf("`%s`.`%s` = ?", d.Statement.Table, col.db), resp.UserId())
				}
			}
		}
	}
}

func applyJoinQuery(d *gorm.DB, resp ScopeResponse) {
	all, scope := resp.DeptScopes()
	applyQueryJoinOn := func(d *gorm.DB, table string, hfs map[string]string, resp ScopeResponse) {
		for tp, name := range hfs {
			switch tp {
			case TypeTenant:
				d = d.Where(fmt.Sprintf("`%s`.`%s` = ?", table, name), resp.TenantId())
			case TypeDept:
				if !all && len(scope) > 0 {
					d = d.Where(fmt.Sprintf("`%s`.`%s` in ?", table, name), scope)
				}
			case TypeUser:
				if !all && len(scope) == 0 {
					d = d.Where(fmt.Sprintf("`%s`.`%s` = ?", table, name), resp.UserId())
				}
			}
		}
		return
	}

	// 处理存在的join查询
	if len(d.Statement.Joins) != 0 {
		for _, join := range d.Statement.Joins {
			if d.Statement.Schema == nil || d.Statement.Schema.Relationships.Relations == nil {
				continue
			}

			relation := d.Statement.Schema.Relationships.Relations[join.Name]
			if relation == nil {
				continue
			}
			hfs := map[string]string{}
			for _, field := range relation.FieldSchema.Fields {
				tv, ok := field.TagSettings[TagHook]
				if ok {
					hfs[tv] = field.DBName
				}
			}
			table := join.Name
			if join.Alias != "" {
				table = join.Alias
			}
			if join.On == nil {
				join.On = &clause.Where{}
			}
			applyQueryJoinOn(d, table, hfs, resp)
		}
	}
}

func applyCondition(d *gorm.DB, resp ScopeResponse) {
	condition := resp.Condition()
	if condition == nil {
		return
	}
	if len(condition.Conditions) == 0 && len(condition.Groups) == 0 {
		return
	}

	var params []interface{}
	whereClause := buildQuery(d.Statement.Table, condition, &params)
	d = d.Where(whereClause, params...)
}

func applySelect(d *gorm.DB, resp ScopeResponse) {
	fields := resp.Fields()
	if len(fields) == 0 {
		return
	}
	for ind, field := range fields {
		if !strings.Contains(field, ".") {
			fields[ind] = fmt.Sprintf("`%s`.`%s`", d.Statement.Table, field)
		}
	}

	if len(d.Statement.Selects) != 0 && d.Statement.Selects[0] == "*" {
		d.Statement.Selects = fields
	}
}

// setValue 设置指定值到指定字段
func checkAndSetValue(value reflect.Value, method, tp string, col column, resp ScopeResponse) error {
	nv, err := getVal(resp, tp)
	if err != nil {
		return err
	}

	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			if err := checkAndSetValue(value.Index(i), method, tp, col, resp); err != nil {
				return err
			}
		}
		return nil
	}

	// 如果是单个对象，则直接设置
	if value.Kind() == reflect.Struct {
		fv := value.FieldByName(col.name)
		if !fv.CanInt() && !fv.CanUint() {
			return errors.New("field is not uint or int")
		}

		// 判断是否存在值，存在则校验值是否在预期之内
		if !fv.IsZero() {
			// 不检查则跳过
			if !col.op[method].check {
				return nil
			}
			ori := cast.ToUint32(fv.Interface())
			return checkVal(resp, tp, ori)
		} else {
			// 不设置则跳过
			if !col.op[method].set {
				return nil
			}
			val := transVal(fv.Kind(), nv)
			fv.Set(reflect.ValueOf(val))
		}
		return nil
	}

	// 如果设置为map,则设置对应的key
	if value.Kind() == reflect.Map {
		mv := value.MapIndex(reflect.ValueOf(col.name))
		if !mv.CanInt() && !mv.CanUint() {
			return errors.New("field is not uint or int")
		}

		// 判断是否存在值，存在则校验值是否在预期之内
		if !mv.IsZero() {
			// 不检查则跳过
			if !col.op[Create].check {
				return nil
			}
			ori := cast.ToUint32(mv.Interface())
			return checkVal(resp, tp, ori)
		} else {
			// 不设置则跳过
			if !col.op[Create].set {
				return nil
			}
			val := transVal(mv.Kind(), nv)
			mv.Set(reflect.ValueOf(val))
		}
		return nil
	}
	return nil
}

func parseTag(tag string) (string, map[string]*op) {
	getVal := func(cu ...bool) map[string]*op {
		defaultOp := func() *op {
			return &op{
				check: false,
				set:   false,
				where: false,
			}
		}
		if len(cu) != 0 && cu[0] {
			return map[string]*op{
				Create: defaultOp(),
				Update: defaultOp(),
				Read:   defaultOp(),
				Delete: defaultOp(),
			}
		}

		return map[string]*op{
			Create: {
				check: true,
				set:   true,
				where: false,
			},
			Update: {
				check: false,
				set:   false,
				where: true,
			},
			Read: {
				check: false,
				set:   false,
				where: true,
			},
			Delete: {
				check: false,
				set:   false,
				where: true,
			},
		}
	}

	getOpKey := func(val string) string {
		switch val {
		case "c":
			return Create
		case "u":
			return Update
		case "r":
			return Read
		case "d":
			return Delete
		default:
			return ""
		}
	}

	setVal := func(method string, value string, obj map[string]*op) {
		ms := strings.Split(method, "")
		vs := strings.Split(value, "")
		vsb := map[string]bool{}
		for _, v := range vs {
			vsb[v] = true
		}
		for _, md := range ms {
			obj[getOpKey(md)].set = vsb["s"]
			obj[getOpKey(md)].check = vsb["c"]
			obj[getOpKey(md)].where = vsb["w"]
		}
	}

	// 提取tag tenant[c:csw]
	// 正则提取[]中的内容
	pv := tagReg.FindString(tag)
	if pv == "" || len(pv) == 2 {
		return tag, getVal()
	}

	key := strings.TrimSuffix(tag, pv)
	// 计算opVal
	opVal := getVal(true)
	pv = pv[1 : len(pv)-1]
	opArr := strings.Split(pv, ",")
	for _, opItem := range opArr {
		opInfo := strings.Split(opItem, ":")
		if len(opInfo) == 2 {
			setVal(opInfo[0], opInfo[1], opVal)
		}
	}

	return key, opVal
}

func checkVal(resp ScopeResponse, tp string, ori uint32) error {
	switch tp {
	case TypeTenant:
		if ori != resp.TenantId() {
			return fmt.Errorf("not tenant scope value is:%d", ori)
		}
		return nil
	case TypeDept:
		all, scope := resp.DeptScopes()
		if !all && !value.InList(scope, ori) {
			return fmt.Errorf("not dept scope value is:%d", ori)
		}
		return nil
	case TypeUser:
		all, scope := resp.DeptScopes()
		if !all && !value.InList(scope, resp.UserDeptId(ori)) {
			return fmt.Errorf("not user scope value is:%d", ori)
		}
		return nil
	default:
		return errors.New("invoke hook type")
	}
}

func getVal(resp ScopeResponse, tp string) (uint32, error) {
	switch tp {
	case TypeTenant:
		return resp.TenantId(), nil
	case TypeDept:
		return resp.DeptId(), nil
	case TypeUser:
		return resp.UserId(), nil
	default:
		return 0, errors.New("invoke hook type")
	}
}

func transVal(kind reflect.Kind, val uint32) any {
	switch kind {
	case reflect.Uint:
		return uint(val)
	case reflect.Uint8:
		return uint8(val)
	case reflect.Uint16:
		return uint16(val)
	case reflect.Uint32:
		return val
	case reflect.Uint64:
		return uint64(val)
	case reflect.Int:
		return int(val)
	case reflect.Int8:
		return int8(val)
	case reflect.Int16:
		return int16(val)
	case reflect.Int32:
		return int32(val)
	case reflect.Int64:
		return int64(val)
	default:
		return val
	}
}

// buildQuery 递归构建SQL查询
func buildQuery(table string, group *ConditionGroup, params *[]interface{}) string {
	var whereClauses []string

	for _, condition := range group.Conditions {
		placeholder := "?"
		whereClause := fmt.Sprintf("`%s`.`%s` %s %s", table, condition.Field, condition.Operator, placeholder)
		whereClauses = append(whereClauses, whereClause)
		*params = append(*params, condition.Value)
	}

	for _, subgroup := range group.Groups {
		subClause := buildQuery(table, subgroup, params)
		if subClause != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("(%s)", subClause))
		}
	}

	if len(whereClauses) > 0 {
		return strings.Join(whereClauses, fmt.Sprintf(" %s ", group.Logic))
	}

	return ""
}
