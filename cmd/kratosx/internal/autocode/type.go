package autocode

type Type struct {
	Keyword string  `json:"keyword"`
	Name    string  `json:"name"`
	Mapping Mapping `json:"-"`
}

type Mapping struct {
	Struct string
	Proto  string
	DB     string
}

const (
	_bool                 = "bool"
	_char                 = "char"
	_string               = "string"
	_varchar64            = "varchar64"
	_varchar128           = "varchar128"
	_varchar256           = "varchar256"
	_shortInteger         = "shortInteger"
	_unsignedShortInteger = "unsignedShortInteger"
	_integer              = "integer"
	_unsignedInteger      = "unsignedInteger"
	_longInteger          = "longInteger"
	_longUnsignedInteger  = "longUnsignedInteger"
	_primaryKey           = "primaryKey"
	_foreignKey           = "foreignKey"
	_time                 = "time"
	_float                = "float"
	_tinytext             = "tinytext"
	_text                 = "text"
	_mediumtext           = "mediumtext"
	_longtext             = "longtext"
	_tinyblob             = "tinyblob"
	_blob                 = "blob"
	_mediumblob           = "mediumblob"
	_longblob             = "longblob"

	_objectTypeList = "list"
	_objectTypeTree = "tree"

	_in      = "in"
	_notIn   = "not in"
	_between = "between"
	_like    = "like"

	_relationHasOne  = "hasOne"
	_relationHasMany = "hasMany"

	_fixedCode = "fixed code"
)

var (
	types = []*Type{
		{
			Keyword: _bool,
			Name:    "布尔值",
			Mapping: Mapping{
				Struct: "bool",
				Proto:  "bool",
				DB:     "bool",
			},
		},
		{
			Keyword: _char,
			Name:    "文本标识",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "char(32) binary",
			},
		},
		{
			Keyword: _varchar64,
			Name:    "短字符",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "varchar(64)",
			},
		},
		{
			Keyword: _varchar128,
			Name:    "中字符",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "varchar(128)",
			},
		},
		{
			Keyword: _varchar256,
			Name:    "长字符",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "varchar(256)",
			},
		},
		{
			Keyword: _shortInteger,
			Name:    "短整数",
			Mapping: Mapping{
				Struct: "int16",
				Proto:  "int16",
				DB:     "tinyint",
			},
		},
		{
			Keyword: _unsignedShortInteger,
			Name:    "短非负数字",
			Mapping: Mapping{
				Struct: "uint16",
				Proto:  "uint16",
				DB:     "unsigned tinyint",
			},
		},
		{
			Keyword: _integer,
			Name:    "整数",
			Mapping: Mapping{
				Struct: "int32",
				Proto:  "int32",
				DB:     "int",
			},
		},
		{
			Keyword: _unsignedInteger,
			Name:    "非负整数",
			Mapping: Mapping{
				Struct: "uint32",
				Proto:  "uint32",
				DB:     "unsigned int",
			},
		},
		{
			Keyword: _longInteger,
			Name:    "长整数",
			Mapping: Mapping{
				Struct: "int64",
				Proto:  "int64",
				DB:     "bigint",
			},
		},
		{
			Keyword: _longUnsignedInteger,
			Name:    "长非负整数",
			Mapping: Mapping{
				Struct: "uint64",
				Proto:  "uint64",
				DB:     "unsigned bigint",
			},
		},
		{
			Keyword: _primaryKey,
			Name:    "主键",
			Mapping: Mapping{
				Struct: "uint32",
				Proto:  "uint32",
				DB:     "unsigned bigint",
			},
		},
		{
			Keyword: _foreignKey,
			Name:    "引用主键",
			Mapping: Mapping{
				Struct: "uint32",
				Proto:  "uint32",
				DB:     "unsigned bigint",
			},
		},
		{
			Keyword: _time,
			Name:    "时间",
			Mapping: Mapping{
				Struct: "int64",
				Proto:  "uint32",
				DB:     "unsigned int",
			},
		},
		{
			Keyword: _float,
			Name:    "浮点数",
			Mapping: Mapping{
				Struct: "float32",
				Proto:  "float",
				DB:     "float",
			},
		},
		{
			Keyword: _tinytext,
			Name:    "短文本",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "tinytext",
			},
		},
		{
			Keyword: _text,
			Name:    "文本",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "text",
			},
		},
		{
			Keyword: _mediumtext,
			Name:    "中长文本",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "mediumtext",
			},
		},
		{
			Keyword: _longtext,
			Name:    "长文本",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "longtext",
			},
		},
		{
			Keyword: _tinyblob,
			Name:    "短二进制",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "tinyblob",
			},
		},
		{
			Keyword: _blob,
			Name:    "二进制",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "blob",
			},
		},
		{
			Keyword: _mediumblob,
			Name:    "中长二进制",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "mediumblob",
			},
		},
		{
			Keyword: _longblob,
			Name:    "长二进制",
			Mapping: Mapping{
				Struct: "string",
				Proto:  "string",
				DB:     "longblob",
			},
		},
	}
)

func Types() []*Type {
	return types
}

func TypesMapping() map[string]Mapping {
	var m = make(map[string]Mapping)
	for _, item := range types {
		m[item.Keyword] = item.Mapping
	}
	return m
}
