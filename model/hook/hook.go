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
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
)

const (
	Read   = "read"
	Create = "create"
	Update = "update"
	Delete = "delete"

	TagHook  = "HOOK"
	TypeDept = "dept"
	TypeUser = "user"
)

type hookContextKey struct{}

type op struct {
	check bool
	set   bool
	where bool
}

type column struct {
	name string
	db   string
	op   map[string]*op
	resp ScopeResponse
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
	// UserDeptId 指定用户的主部门
	UserDeptId(uint32) uint32
	// Condition 条件约束
	Condition() *ConditionGroup
	// Fields 查询字段
	Fields() []string
}

var tagReg = regexp.MustCompile(`\[(.*?)\]`)

func Apply(d *gorm.DB, dbName, method string, req ScopeRequestFunc) {
	originCtx := d.Statement.Context
	if originCtx == nil {
		originCtx = context.Background()
	}
	_, is := originCtx.Value(hookContextKey{}).(struct{})
	if is {
		return
	}

	d.Statement.Context = context.WithValue(originCtx, hookContextKey{}, struct{}{})
	defer func() {
		d.Statement.Context = originCtx
	}()

	if req == nil {
		return
	}

	// 获取当前表权限数据
	curResp, err := req(d.Statement.Context, dbName, d.Statement.Table, method)
	if err != nil {
		_ = d.AddError(err)
		return
	}
	if curResp == nil {
		return
	}

	hfs := map[string][]column{}
	if d.Statement.Schema != nil {
		for _, field := range d.Statement.Schema.Fields {
			tv, ok := field.TagSettings[TagHook]
			key, opv, table := parseTag(tv)
			if table == "" {
				table = d.Statement.Table
			}
			if ok {
				// 获取权限数据
				resp, err := req(d.Statement.Context, dbName, table, method)
				if err != nil {
					_ = d.AddError(err)
					return
				}
				if resp == nil {
					continue
				}

				hfs[key] = append(hfs[key], column{
					name: field.Name,
					db:   field.DBName,
					op:   opv,
					resp: resp,
				})
			}
		}
	}

	switch method {
	case Create:
		applyCheckAndSet(d, Create, hfs)
	case Update:
		applyCheckAndSet(d, Update, hfs)
		applyWhere(d, Update, hfs)
		applyCondition(d, curResp)
	case Read:
		applySelect(d, curResp)
		applyWhere(d, Read, hfs)
		applyCondition(d, curResp)
		applyJoinQuery(d, dbName, method, req)
	case Delete:
		applyWhere(d, Delete, hfs)
		applyCondition(d, curResp)
	}
}

func applyCheckAndSet(d *gorm.DB, method string, hfs map[string][]column) {
	for tp, cols := range hfs {
		var err error
		for _, col := range cols {
			err = checkAndSetValue(d.Statement.ReflectValue, method, tp, col, col.resp)
			if err != nil {
				_ = d.AddError(err)
				return
			}
		}

	}
}

func applyWhere(d *gorm.DB, method string, hfs map[string][]column) {

	for tp, cols := range hfs {
		for _, col := range cols {
			if col.op == nil || col.op[method] == nil {
				continue
			}
			resp := col.resp
			all, scope := resp.DeptScopes()
			if !col.op[method].where {
				continue
			}

			switch tp {
			case TypeDept:
				if !all {
					if len(scope) == 0 {
						d = d.Where("1 = 0")
					} else {
						d = d.Where(clause.IN{
							Column: clause.Column{Table: d.Statement.Table, Name: col.db},
							Values: uint32Values(scope),
						})
					}
				}
			case TypeUser:
				if !all && len(scope) == 0 {
					d = d.Where(clause.Eq{
						Column: clause.Column{Table: d.Statement.Table, Name: col.db},
						Value:  resp.UserId(),
					})
				}
			}
		}
	}
}

func applyJoinQuery(d *gorm.DB, database string, method string, req ScopeRequestFunc) {

	applyQueryJoinOn := func(table string, hfs map[string]string, resp ScopeResponse) []clause.Expression {
		all, scope := resp.DeptScopes()
		var exprs []clause.Expression
		for tp, name := range hfs {
			switch tp {
			case TypeDept:
				if !all {
					if len(scope) == 0 {
						exprs = append(exprs, clause.Expr{SQL: "1 = 0"})
					} else {
						exprs = append(exprs, clause.IN{
							Column: clause.Column{Table: table, Name: name},
							Values: uint32Values(scope),
						})
					}
				}
			case TypeUser:
				if !all && len(scope) == 0 {
					exprs = append(exprs, clause.Eq{
						Column: clause.Column{Table: table, Name: name},
						Value:  resp.UserId(),
					})
				}
			}
		}
		return exprs
	}

	// 处理存在的join查询
	if len(d.Statement.Joins) != 0 {
		originJoins := append(d.Statement.Joins[:0:0], d.Statement.Joins...)
		d.Statement.Joins = d.Statement.Joins[:0]
		for _, join := range originJoins {
			if d.Statement.Schema == nil || d.Statement.Schema.Relationships.Relations == nil {
				d.Statement.Joins = append(d.Statement.Joins, join)
				continue
			}

			relations := joinRelations(d.Statement.Schema.Relationships.Relations, join.Name)
			if len(relations) == 0 {
				d.Statement.Joins = append(d.Statement.Joins, join)
				continue
			}

			parentTableName := clause.CurrentTable
			for idx, relation := range relations {
				tableAliasName := relation.Name
				if parentTableName != clause.CurrentTable {
					tableAliasName = utils.NestedRelationName(parentTableName, tableAliasName)
				}
				if idx == len(relations)-1 && join.Alias != "" {
					tableAliasName = join.Alias
				}

				// 获取权限数据
				resp, err := req(d.Statement.Context, database, relation.FieldSchema.Table, method)
				if err != nil {
					_ = d.AddError(err)
					return
				}

				hfs := map[string]string{}
				for _, field := range relation.FieldSchema.Fields {
					tv, ok := field.TagSettings[TagHook]
					if ok {
						key, _, _ := parseTag(tv)
						hfs[key] = field.DBName
					}
				}
				exprs := applyQueryJoinOn(tableAliasName, hfs, resp)
				if len(relations) == 1 {
					if len(exprs) != 0 {
						if join.On == nil {
							join.On = &clause.Where{}
						}
						join.On.Exprs = append(join.On.Exprs, exprs...)
					}
					d.Statement.Joins = append(d.Statement.Joins, join)
				} else {
					nestedJoin := join
					nestedJoin.Name = joinPath(relations, idx)
					nestedJoin.Alias = ""
					nestedJoin.Selects = nil
					nestedJoin.Omits = nil
					if len(exprs) != 0 {
						nestedJoin.On = &clause.Where{Exprs: exprs}
					} else {
						nestedJoin.On = nil
					}
					d.Statement.Joins = append(d.Statement.Joins, nestedJoin)
				}

				parentTableName = tableAliasName
			}
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
	whereClause := buildQuery(d.Statement, d.Statement.Table, condition, &params)
	if whereClause == "" {
		return
	}
	d = d.Where(whereClause, params...)
}

func applySelect(d *gorm.DB, resp ScopeResponse) {
	fields := resp.Fields()
	if len(fields) == 0 {
		return
	}
	selects := make([]clause.Column, 0, len(fields))
	for _, field := range fields {
		if col, ok := selectColumn(d.Statement.Table, field); ok {
			selects = append(selects, col)
		}
	}
	if len(selects) == 0 {
		return
	}
	d.Statement.Selects = nil
	d.Statement.AddClause(clause.Select{Distinct: d.Statement.Distinct, Columns: selects})
}

// setValue 设置指定值到指定字段
func checkAndSetValue(value reflect.Value, method, tp string, col column, resp ScopeResponse) error {
	nv, err := getVal(resp, tp)
	if err != nil {
		return err
	}
	if col.op == nil || col.op[method] == nil {
		return nil
	}
	opv := col.op[method]

	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil
		}
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

		if fv.Kind() == reflect.Ptr {
			if !fv.IsNil() {
				// 不检查则跳过
				if !opv.check {
					return nil
				}
				ori := cast.ToUint32(fv.Elem().Interface())
				return checkVal(resp, tp, ori)
			} else {
				// 不设置则跳过
				if !opv.set {
					return nil
				}
				ptr := reflect.New(fv.Type().Elem())
				if err := setReflectInt(ptr.Elem(), nv); err != nil {
					return err
				}
				fv.Set(ptr)
			}

		} else {
			if !fv.CanInt() && !fv.CanUint() {
				return errors.New("field is not uint or int")
			}

			// 判断是否存在值，存在则校验值是否在预期之内
			if !fv.IsZero() {
				// 不检查则跳过
				if !opv.check {
					return nil
				}
				ori := cast.ToUint32(fv.Interface())
				return checkVal(resp, tp, ori)
			} else {
				// 不设置则跳过
				if !opv.set {
					return nil
				}
				val := transVal(fv.Kind(), nv)
				fv.Set(reflect.ValueOf(val))
			}
		}
		return nil
	}

	// 如果设置为map,则设置对应的key
	if value.Kind() == reflect.Map {
		return checkAndSetMapValue(value, method, tp, col, resp, nv)
	}
	return nil
}

func checkAndSetMapValue(value reflect.Value, method, tp string, col column, resp ScopeResponse, nv uint32) error {
	if value.IsNil() {
		return nil
	}
	if col.op == nil || col.op[method] == nil {
		return nil
	}

	key, err := mapHookKey(value, col)
	if err != nil {
		return err
	}

	mv := value.MapIndex(key)
	if !mv.IsValid() {
		if !col.op[method].set {
			return nil
		}
		return setMapValue(value, key, nv)
	}

	if mv.Kind() == reflect.Interface {
		if mv.IsNil() {
			if !col.op[method].set {
				return nil
			}
			return setMapValue(value, key, nv)
		}
		mv = mv.Elem()
	}

	if mv.Kind() == reflect.Ptr {
		if !mv.IsNil() {
			if !col.op[method].check {
				return nil
			}
			ori := cast.ToUint32(mv.Elem().Interface())
			return checkVal(resp, tp, ori)
		}
		if !col.op[method].set {
			return nil
		}
		ptr := reflect.New(mv.Type().Elem())
		if err := setReflectInt(ptr.Elem(), nv); err != nil {
			return err
		}
		value.SetMapIndex(key, ptr)
		return nil
	}

	if !mv.CanInt() && !mv.CanUint() {
		return errors.New("field is not uint or int")
	}

	if !mv.IsZero() {
		if !col.op[method].check {
			return nil
		}
		ori := cast.ToUint32(mv.Interface())
		return checkVal(resp, tp, ori)
	}
	if !col.op[method].set {
		return nil
	}
	return setMapValue(value, key, nv)
}

func setMapValue(value, key reflect.Value, nv uint32) error {
	elemType := value.Type().Elem()
	if elemType.Kind() == reflect.Interface {
		value.SetMapIndex(key, reflect.ValueOf(nv))
		return nil
	}
	if elemType.Kind() == reflect.Ptr {
		ptr := reflect.New(elemType.Elem())
		if err := setReflectInt(ptr.Elem(), nv); err != nil {
			return err
		}
		value.SetMapIndex(key, ptr)
		return nil
	}
	if !isIntKind(elemType.Kind()) {
		return errors.New("field is not uint or int")
	}
	val := reflect.ValueOf(transVal(elemType.Kind(), nv))
	if val.Type().AssignableTo(elemType) {
		value.SetMapIndex(key, val)
		return nil
	}
	if val.Type().ConvertibleTo(elemType) {
		value.SetMapIndex(key, val.Convert(elemType))
		return nil
	}
	return errors.New("field is not uint or int")
}

func mapHookKey(value reflect.Value, col column) (reflect.Value, error) {
	for _, name := range []string{col.name, col.db} {
		key, err := mapKey(value, name)
		if err != nil {
			return reflect.Value{}, err
		}
		if key.IsValid() && value.MapIndex(key).IsValid() {
			return key, nil
		}
	}

	name := col.db
	if name == "" {
		name = col.name
	}
	return mapKey(value, name)
}

func mapKey(value reflect.Value, name string) (reflect.Value, error) {
	if name == "" {
		return reflect.Value{}, nil
	}
	key := reflect.ValueOf(name)
	if key.Type().AssignableTo(value.Type().Key()) {
		return key, nil
	}
	if key.Type().ConvertibleTo(value.Type().Key()) {
		return key.Convert(value.Type().Key()), nil
	}
	return reflect.Value{}, errors.New("map key is not string")
}

func setReflectInt(value reflect.Value, nv uint32) error {
	if value.CanUint() {
		value.SetUint(uint64(nv))
		return nil
	}
	if value.CanInt() {
		value.SetInt(int64(nv))
		return nil
	}
	return errors.New("field is not uint or int")
}

func isIntKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func uint32Values(values []uint32) []interface{} {
	res := make([]interface{}, 0, len(values))
	for _, item := range values {
		res = append(res, item)
	}
	return res
}

func selectColumn(table string, field string) (clause.Column, bool) {
	field = strings.TrimSpace(field)
	if field == "" {
		return clause.Column{}, false
	}
	parts := strings.Split(field, ".")
	if len(parts) == 1 {
		return clause.Column{Table: table, Name: parts[0]}, true
	}
	if len(parts) == 2 && isValidColumnName(parts[0]) && isValidColumnName(parts[1]) {
		return clause.Column{Table: parts[0], Name: parts[1]}, true
	}
	return clause.Column{}, false
}

func quotedColumn(stmt *gorm.Statement, table string, field string) string {
	return stmt.Quote(clause.Column{Table: table, Name: field})
}

func joinRelations(relations map[string]*schema.Relationship, joinName string) []*schema.Relationship {
	relation, ok := relations[joinName]
	if ok {
		return []*schema.Relationship{relation}
	}

	names := strings.Split(joinName, ".")
	if len(names) <= 1 {
		return nil
	}

	currentRelations := relations
	res := make([]*schema.Relationship, 0, len(names))
	for _, name := range names {
		relation, ok = currentRelations[name]
		if !ok {
			return nil
		}
		res = append(res, relation)
		currentRelations = relation.FieldSchema.Relationships.Relations
	}
	return res
}

func joinPath(relations []*schema.Relationship, idx int) string {
	names := make([]string, 0, idx+1)
	for i := 0; i <= idx; i++ {
		names = append(names, relations[i].Name)
	}
	return strings.Join(names, ".")
}

func parseTag(tag string) (string, map[string]*op, string) {
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
			opKey := getOpKey(md)
			if opKey == "" || obj[opKey] == nil {
				continue
			}
			obj[opKey].set = vsb["s"]
			obj[opKey].check = vsb["c"]
			obj[opKey].where = vsb["w"]
		}
	}

	// 提取tag tenant[c:csw]
	// 正则提取[]中的内容
	pv := tagReg.FindString(tag)
	key := strings.TrimSuffix(tag, pv)
	if pv == "" || len(pv) == 2 {
		return key, getVal(), ""
	}

	// 计算opVal
	opVal := getVal(true)
	pv = pv[1 : len(pv)-1]

	opInfo := strings.Split(pv, ":")
	if len(opInfo) == 1 {
		return key, getVal(), opInfo[0]
	}
	if len(opInfo) == 2 {
		setVal(opInfo[0], opInfo[1], opVal)
		return key, opVal, ""
	}
	if len(opInfo) == 3 {
		setVal(opInfo[1], opInfo[2], opVal)
		return key, opVal, opInfo[0]
	}

	return "", nil, ""
}

func checkVal(resp ScopeResponse, tp string, ori uint32) error {
	switch tp {
	case TypeDept:
		all, scope := resp.DeptScopes()
		if !all && !value.InList(scope, ori) {
			return fmt.Errorf("not dept scope value is %d", ori)
		}
		return nil
	case TypeUser:
		all, scope := resp.DeptScopes()
		if !all && !value.InList(scope, resp.UserDeptId(ori)) {
			return fmt.Errorf("not user scope value is %d", ori)
		}
		return nil
	default:
		return errors.New("invoke hook type")
	}
}

func getVal(resp ScopeResponse, tp string) (uint32, error) {
	switch tp {
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
func buildQuery(stmt *gorm.Statement, table string, group *ConditionGroup, params *[]interface{}) string {
	var whereClauses []string

	logic := strings.ToUpper(group.Logic)
	if logic != "AND" && logic != "OR" {
		logic = "AND"
	}

	for _, condition := range group.Conditions {
		if condition == nil {
			continue
		}
		if !isValidColumnName(condition.Field) {
			continue
		}
		if !isValidOperator(condition.Operator) {
			continue
		}
		operator := strings.ToUpper(strings.TrimSpace(condition.Operator))
		column := quotedColumn(stmt, table, condition.Field)
		whereClause := fmt.Sprintf("%s %s ?", column, operator)
		if operator == "IN" || operator == "NOT IN" {
			whereClause = fmt.Sprintf("%s %s (?)", column, operator)
		}
		if operator == "IS NULL" || operator == "IS NOT NULL" {
			whereClause = fmt.Sprintf("%s %s", column, operator)
			whereClauses = append(whereClauses, whereClause)
			continue
		}
		whereClauses = append(whereClauses, whereClause)
		*params = append(*params, condition.Value)
	}

	for _, subgroup := range group.Groups {
		if subgroup == nil {
			continue
		}
		subClause := buildQuery(stmt, table, subgroup, params)
		if subClause != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("(%s)", subClause))
		}
	}

	if len(whereClauses) > 0 {
		return strings.Join(whereClauses, fmt.Sprintf(" %s ", logic))
	}

	return ""
}

var validOperators = map[string]bool{
	"=": true, "!=": true, "<>": true,
	">": true, "<": true, ">=": true, "<=": true,
	"LIKE": true, "NOT LIKE": true,
	"IN": true, "NOT IN": true,
	"IS NULL": true, "IS NOT NULL": true,
}

func isValidOperator(op string) bool {
	return validOperators[strings.ToUpper(strings.TrimSpace(op))]
}

func isValidColumnName(name string) bool {
	if name == "" {
		return false
	}
	for _, c := range name {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
			return false
		}
	}
	return true
}
