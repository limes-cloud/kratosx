package config

import (
	"time"
)

// App 配置结构体
type App struct {
	ID      string    // 唯一ID
	Name    string    // 应用名称
	Version string    // 应用版本
	Env     string    // 环境
	Server  *struct { // 服务配置
		Count    int          // 服务数量
		Registry *string      // 服务注册中心
		Http     *HttpService // HTTP服务
		Grpc     *GrpcService // GRPC服务
		TLS      *TLS         // TLS配置
	}
	Signature      *Signature          // 签名配置
	Logger         *Logger             // 日志配置
	Pool           *Pool               // 连接池配置
	Email          *Email              // 邮件配置
	JWT            *JWT                // JWT配置
	Request        *Request            // http request配置
	Logging        *Logging            // 日志记录配置
	Authentication *Authentication     // 认证配置
	Tracing        *Tracing            // 链路追踪配置
	Client         []*Client           // 客户端配置
	Database       []*Database         // 数据库配置
	Redis          []*Redis            // Redis配置
	Loader         map[string]string   // 加载器配置
	Captcha        map[string]*Captcha // 验证码配置
	RateLimit      bool                // 限流配置
	Metrics        bool                // 指标配置
	Prometheus     []*Prometheus       // Prometheus配置
}

type TLS struct {
	Name string
	Ca   string
	Pem  string
	Key  string
}

type GrpcService struct {
	Network        string
	Host           string
	Port           int
	MaxRecvSize    int
	Timeout        time.Duration
	TimeoutSpecial map[string]time.Duration
}

type HttpService struct {
	Network        string
	Host           string
	Port           int
	Timeout        time.Duration
	TimeoutSpecial map[string]time.Duration
	FormatResponse bool
	Cors           *Cors
	Marshal        *Marshal
	Pprof          *Pprof
	WebServerDir   string
}

type Pprof struct {
	Query  string
	Secret string
}

type Database struct {
	Enable     bool
	Name       string
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
	Size             int           // 最大协程数量
	ExpiryDuration   time.Duration // 过期时间
	PreAlloc         bool          // 是否预分配
	MaxBlockingTasks int           // 最大的并发任务
	Nonblocking      bool          // 设置为true时maxBlockingTasks将失效，不限制并发任务
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
	Length       int
	Expire       time.Duration
	Redis        string
	RefreshTime  time.Duration
	Limit        int
	UniqueDevice bool
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

type LoggerFile struct {
	ErrorAlone bool          `json:"errorAlone"` // 错误日志单独输出
	Name       string        `json:"name"`       // 日志文件名
	SplitTime  time.Duration `json:"splitTime"`  // 日志切割时间间隔
	MaxAge     time.Duration `json:"maxAge"`     // 备份文件最大保存时间
	MaxBackup  int           `json:"maxBackup"`  // 备份文件最大个数，MaxAge可能仍会导致它们被删除。
}

type Logger struct {
	Level      int8        `json:"level"`      // 日志等级 0：info 1:warning 2:error
	Output     []string    `json:"output"`     // 输出位置，支持 stdout,file
	EnCoder    string      `json:"enCoder"`    // 编码器类型 json,console
	Caller     bool        `json:"caller"`     // 显示调用者信息，默认不显示
	CallerSkip int         `json:"callerSkip"` // 调用者层级，默认 2
	HookStd    bool        `json:"hookStd"`    // 开启劫持标准输出
	File       *LoggerFile `json:"file"`       // 日志文件配置
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

type Request struct {
	EnableLog        bool          `json:"enableLog"`        // 是否开启请求日志
	RetryCount       int           `json:"retryCount"`       // 最大重试次数
	RetryWaitTime    time.Duration `json:"retryWaitTime"`    // 重试等待时间
	MaxRetryWaitTime time.Duration `json:"maxRetryWaitTime"` // 最大重试等待时间
	Timeout          time.Duration `json:"timeout"`          // 请求超时时间
	UserAgent        string        `json:"userAgent"`        // 请求服务名称，UA
}

type Client struct {
	Server    string
	Type      string
	Timeout   time.Duration
	Metadata  map[string]string
	Backends  []Backend
	Signature *Signature
	TLS       *TLS
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
