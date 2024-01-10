package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"m-sec/common/logs"
	"path"
	"runtime"
)

var Conf = InitConfig()

type Config struct {
	viper        *viper.Viper
	ServerConfig *ServerConfig
	GrpcConfig   *GrpcConfig
	EtcdConfig   *EtcdConfig
	MysqlConfig  *MysqlConfig
	JwtConfig    *JwtConfig
	LimitConfig  *LimitConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name    string
	Addr    string
	Version string
	Weight  int64
}

type EtcdConfig struct {
	Addrs []string
}

type MysqlConfig struct {
	Username     string
	Password     string
	Host         string
	Port         int
	Db           string
	MaxOpenConns int
	MaxIdleConns int
}

type JwtConfig struct {
	AccessExp     int
	RefreshExp    int
	AccessSecret  string
	RefreshSecret string
}

type LimitConfig struct {
	UrlQPS int
	IpQPS  int
}

func InitConfig() *Config {
	conf := &Config{viper: viper.New()}
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	// 此路径仅在开发过程中调试使用
	conf.viper.AddConfigPath(getCurrentAbPathByCaller())
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	conf.InitServerConfig()
	conf.InitZapLog()
	conf.InitMysqlConfig()
	conf.InitJwtConfig()
	conf.InitLimitConfig()

	return conf
}

func (c *Config) InitServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.ServerConfig = sc
}

func (c *Config) InitZapLog() {
	// 从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) ReadRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("cache.host") + ":" + c.viper.GetString("cache.port"),
		Password: c.viper.GetString("cache.password"), // no password set
		DB:       c.viper.GetInt("cache.db"),          // use default DB
	}
}

func (c *Config) InitMysqlConfig() {
	mc := &MysqlConfig{
		Username:     c.viper.GetString("mysql.username"),
		Password:     c.viper.GetString("mysql.password"),
		Host:         c.viper.GetString("mysql.host"),
		Port:         c.viper.GetInt("mysql.port"),
		Db:           c.viper.GetString("mysql.db"),
		MaxOpenConns: c.viper.GetInt("mysql.maxOpenConns"),
		MaxIdleConns: c.viper.GetInt("mysql.maxIdleConns"),
	}
	c.MysqlConfig = mc
}

func (c *Config) InitJwtConfig() {
	jc := &JwtConfig{
		AccessExp:     c.viper.GetInt("jwt.accessExp"),
		RefreshExp:    c.viper.GetInt("jwt.refreshExp"),
		AccessSecret:  c.viper.GetString("jwt.accessSecret"),
		RefreshSecret: c.viper.GetString("jwt.refreshSecret"),
	}
	c.JwtConfig = jc
}

func (c *Config) InitLimitConfig() {
	lc := &LimitConfig{
		UrlQPS: c.viper.GetInt("limit.urlQps"),
		IpQPS:  c.viper.GetInt("limit.ipQps"),
	}
	c.LimitConfig = lc
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
