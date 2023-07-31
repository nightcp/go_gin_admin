package core

import (
	"flag"
	"github.com/spf13/viper"
)

var Config = initConfig(".")

type envConfig struct {
	AppName                string   `mapstructure:"APP_NAME"`              // 项目名称
	JwtKey                 string   `mapstructure:"JWT_KEY"`               // JWT密钥
	JwtTTl                 int      `mapstructure:"JWT_TTL"`               // Jwt Token有效时长, 单位: 分钟
	JwtRenewTTl            int      `mapstructure:"JWT_RENEW_TTL"`         // Jwt Token续期时长, 单位: 分钟
	GinMode                string   `mapstructure:"GIN_MODE"`              // 模式
	ServerPort             string   `mapstructure:"SERVER_PORT"`           // 服务端口
	DBDSN                  string   `mapstructure:"DB_DSN"`                // DB数据源名
	RedisPrefix            string   `mapstructure:"REDIS_PREDIX"`          // Redis前缀
	RedisUrl               string   `mapstructure:"REDIS_URL"`             // Redis连接字符串
	LogPath                string   `mapstructure:"LOG_PATH"`              // 日志文件保存目录
	LogMaxSize             int      `mapstructure:"LOG_MAX_SIZE"`          // 在进行切割之前, 日志文件的最大大小, 单位: MB
	LogMaxBackups          int      `mapstructure:"LOG_MAX_BACKUPS"`       // 保留旧文件的最大个数
	LogMaxAge              int      `mapstructure:"LOG_MAX_AGE"`           // 保留旧文件的最大天数
	LogCompress            bool     `mapstructure:"LOG_COMPRESS"`          // 是否压缩/归档旧文件
	RateLimiterCapacity    int64    `mapstructure:"RATE_LIMITER_CAPACITY"` // 限流最大令牌数
	RateLimiterQuantum     int64    `mapstructure:"RATE_LIMITER_QUANTUM"`  // 限流每秒补充令牌数
	DBMaxIdleConns         int      // 数据库空闲连接池最大值
	DBMaxOpenConns         int      // 数据库连接池最大值
	DBConnMaxLifetimeHours int16    // 连接可复用的最大时间(小时)
	RedisPoolSize          int      // Redis连接池大小
	ReadTimeout            int      // 服务最大读取时间(秒)
	WriteTimeout           int      // 服务最大写入时间(秒)
	Locale                 string   // 默认地区语言
	UploadFilePath         string   // 上传文件路径
	UploadFileType         []string // 上传文件限制类型
	UploadFileMaxSize      int64    // 上传文件限制大小, 单位: m
	FileRequestPath        string   // 文件请求相对路径
}

// initConfig 初始化配置
func initConfig(path string) envConfig {
	var cfgPath string
	flag.StringVar(&cfgPath, "c", "", "config file path.")
	flag.Parse()
	if cfgPath == "" {
		viper.AddConfigPath(path)
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(cfgPath)
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic("InitConfig ReadInConfig err: " + err.Error())
	}
	config := envConfig{
		DBMaxIdleConns:         10,
		DBMaxOpenConns:         100,
		DBConnMaxLifetimeHours: 2,
		RedisPoolSize:          100,
		ReadTimeout:            10,
		WriteTimeout:           10,
		Locale:                 "zh_CN",
		UploadFilePath:         "storage/uploads/",
		UploadFileType:         []string{".png", ".jpg", ".jpeg"},
		UploadFileMaxSize:      2,
		FileRequestPath:        "storage",
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic("InitConfig Unmarshal err: " + err.Error())
	}
	return config
}
