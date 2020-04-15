package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lexkong/log"
)

// C 配置文件全局变量
var c *Conf

func init() {
	// 加载.env文件
	err := godotenv.Load()
	if err != nil {
		log.Info("Error loading .env file")
	}

	c = initConf()
}

// Conf 全局配置参数
type Conf struct {
	// 应用名称
	appName string

	keyDelim string

	port string

	// debug
	appDebug bool

	// 时区
	timezone string

	// 默认连接数据库
	connectDefault string

	// 数据库连接
	connections map[string]SQLConf

	// Redis
	redis map[string]Redis
}

// SQLConf 参数
type SQLConf struct {
	driver   string
	host     string
	port     string
	database string
	username string
	password string
	charset  string
}

// Redis 参数
type Redis struct {
	host     string
	port     string
	password string
	database string
}

// 初始化配置
func initConf() *Conf {
	c := new(Conf)
	c.appName = envStr("APP_NAME", "Go")
	c.port = envStr("PORT", "8080")
	c.appDebug = envBool("APP_DEBUG", false)
	c.timezone = envStr("TIMEZONE", "PRC")
	c.connectDefault = "mysql"
	c.connections = map[string]SQLConf{
		"mysql": SQLConf{
			"mysql",
			envStr("MYSQL_HOST", "127.0.0.1"),
			envStr("MYSQL_PORT", "3306"),
			envStr("MYSQL_DATABASE", ""),
			envStr("MYSQL_USERNAME", "root"),
			envStr("MYSQL_PASSWORD", ""),
			"utf8",
		},
	}
	c.redis = map[string]Redis{
		"redis": Redis{
			envStr("REDIS_HOST", "127.0.0.1"),
			envStr("REDIS_HOST", "6379"),
			envStr("REDIS_PASSWORD", ""),
			envStr("REDIS_DATABASE", ""),
		},
	}
	return c
}

// 参数设置
func envStr(key string, def string) string {
	upperKey := strings.ToUpper(key)
	if option, ok := os.LookupEnv(upperKey); ok {
		return option
	}
	return def
}

func envBool(key string, def bool) bool {
	upperKey := strings.ToUpper(key)
	if option, ok := os.LookupEnv(upperKey); ok {
		tran, err := strconv.ParseBool(option)
		if err != nil {
			return false
		}
		return tran
	}
	return def

}

// GetStr 获取.env变量String型
func GetStr(key string) string { return c.GetStr(key) }

// GetStr 类方法
func (c *Conf) GetStr(key string) string {
	upperKey := strings.ToUpper(key)
	return os.Getenv(upperKey)
}

// GetInt 获取.env变量Int型
func GetInt(key string) int { return c.GetInt(key) }

// GetInt 类方法
func (c *Conf) GetInt(key string) int {
	upperKey := strings.ToUpper(key)
	if option, ok := os.LookupEnv(upperKey); ok {
		tran, err := strconv.ParseInt(option, 10, 0)
		if err != nil {
			return 0
		}
		return int(tran)
	}
	return 0
}

// func getCurrentDirectory() string {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	return strings.Replace(dir, "\\", "/", -1)
// }
