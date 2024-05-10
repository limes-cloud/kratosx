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
		Count    int
		Registry *string
		Http     *HttpService
		Grpc     *GrpcService
		Tls      *Tls
	}
	Signature      *Signature
	Log            *Logger
	Pool           *Pool
	Email          *Email
	JWT            *JWT
	Http           *Http
	Logging        *Logging
	Authentication *Authentication
	Tracing        *Tracing
	Client         []*Client
	Database       map[string]*Database
	Redis          map[string]*Redis
	Loader         map[string]string
	Captcha        map[string]*Captcha
	RateLimit      bool
	Metrics        bool
	Prometheus     []*Prometheus
}

type Tls struct {
	Name string
	Ca   string
	Pem  string
	Key  string
}

type GrpcService struct {
	Network     string
	Addr        string
	Host        string
	Port        int
	MaxRecvSize int
	Timeout     time.Duration
}

type HttpService struct {
	Network        string
	Addr           string
	Host           string
	Port           int
	Timeout        time.Duration
	FormatResponse bool
	Cors           *Cors
	Marshal        *Marshal
	Pprof          *Pprof
}

type Pprof struct {
	Query  string
	Secret string
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
	Initializer    *DBInitializer
}

type DBTransformError struct {
	Enable bool
	Format *struct {
		Duplicated *string
		AddForeign *string
		DelForeign *string
	}
}

type DBInitializer struct {
	Enable bool
	Force  bool
	Path   string
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
	Refresh  bool
	IpLimit  int
	Template string
}

type Logging struct {
	Enable    bool
	Whitelist map[string]bool
}

type JWT struct {
	EnableGrpc bool
	Redis      string
	Header     string
	Secret     string
	Unique     bool
	UniqueKey  string
	Expire     time.Duration
	Renewal    time.Duration
	Whitelist  map[string]bool
}

type Authentication struct {
	EnableGrpc bool
	DB         string
	Prefix     string
	Redis      string
	RoleKey    string
	Whitelist  map[string]bool
	SkipRole   []string
}

type Logger struct {
	Level      int8
	Output     []string
	EnCoder    string
	Caller     bool
	CallerSkip *int
	File       *struct {
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

type Signature struct {
	Enable    bool
	Ak        string
	Sk        string
	Whitelist map[string]bool
	Expire    time.Duration
}

type Http struct {
	EnableLog        bool
	RetryCount       int
	RetryWaitTime    time.Duration
	MaxRetryWaitTime time.Duration
	Timeout          time.Duration
}

type Client struct {
	Server    string
	Type      string
	Timeout   time.Duration
	Metadata  map[string]string
	Backends  []Backend
	Signature *Signature
	Tls       *Tls
}

type Backend struct {
	Target string
	Weight *int64
}

type Prometheus struct {
	Name       string
	Type       string
	Help       string
	Namespace  string
	Subsystem  string
	Buckets    []float64
	Labels     []string
	Objectives map[float64]float64
	MaxAge     time.Duration
	AgeBuckets uint32
	BufCap     uint32
}
