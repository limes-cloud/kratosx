package config

import (
	"time"
)

type App struct {
	ID      string
	Name    string
	Version string
	Env     string
	Server  *struct {
		Http *HttpService
		Grpc *GrpcService
	}
	Log            *Logger
	Pool           *Pool
	Email          *Email
	JWT            *JWT
	Logging        *Logging
	Authentication *Authentication
	Tracing        *Tracing
	Database       map[string]*Database
	Redis          map[string]*Redis
	Loader         map[string]string
	Captcha        map[string]*Captcha
	RateLimit      bool
}

type GrpcService struct {
	Network     string
	Addr        string
	MaxRecvSize int
	Timeout     time.Duration
}

type HttpService struct {
	Network        string
	Addr           string
	Timeout        time.Duration
	FormatResponse bool
	Cors           *Cors
	Marshal        *Marshal
}

type Database struct {
	Enable     bool
	Drive      string
	AutoCreate bool
	Connect    DBConnect
	Config     DBConfig
}

type DBConnect struct {
	Username string
	Password string
	Host     string
	Port     int
	Option   string
	DBName   string
}

type DBConfig struct {
	TablePrefix    string
	Connect        DBConnect
	LogLevel       int
	PrepareStmt    bool
	DryRun         bool
	TransformError *DBTransformError
	SlowThreshold  time.Duration
	MaxLifetime    time.Duration
	MaxOpenConn    int
	MaxIdleConn    int
}

type DBTransformError struct {
	Enable bool
	Format *struct {
		Duplicated *string
		AddForeign *string
		DelForeign *string
	}
}

type Redis struct {
	Enable   bool
	Name     string
	Host     string
	Username string
	Password string
}

type Pool struct {
	Size             int
	ExpiryDuration   time.Duration
	PreAlloc         bool
	MaxBlockingTasks int
	Nonblocking      bool
}

type Email struct {
	User     string
	Name     string
	Password string
	Host     string
	Port     int
	Template map[string]EmailTemplate
}

type EmailTemplate struct {
	Subject string
	Path    string
	Enable  *bool
	From    string
	Type    string
}

type Captcha struct {
	Type     string
	Length   int
	Expire   time.Duration
	Redis    string
	Height   int
	Width    int
	Skew     float64
	DotCount int
	Template string
}

type Logging struct {
	Enable    bool
	Whitelist map[string]bool
}

type JWT struct {
	Redis     string
	Header    string
	Secret    string
	Expire    time.Duration
	Renewal   time.Duration
	Whitelist map[string]bool
}

type Authentication struct {
	DB        string
	Prefix    string
	Redis     string
	RoleKey   string
	Whitelist map[string]bool
	SkipRole  []string
}

type Logger struct {
	Level  int8
	Output []string
	File   *struct {
		Name      string
		MaxSize   int
		MaxBackup int
		MaxAge    int
		Compress  bool
		LocalTime bool
	}
}

type Tracing struct {
	HttpEndpoint string
	SampleRatio  *float32
	Timeout      time.Duration
	Insecure     *bool
}

type Cors struct {
	AllowCredentials    bool
	AllowOrigins        []string
	AllowMethods        []string
	AllowHeaders        []string
	ExposeHeaders       []string
	MaxAge              time.Duration
	AllowPrivateNetwork bool
}

type Marshal struct {
	ForceUseJson    bool
	EmitUnpopulated bool
	UseProtoNames   bool
}
