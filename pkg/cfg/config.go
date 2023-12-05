package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type TomlConfig struct {
	Log    LogConfig
	Path   PathConfig
	Server ServerConfig
	Db     DbConfig
}
type ServerConfig struct {
	Addr string
	Port int
	Cors string
}

// 日志保存地址
type LogConfig struct {
	Path  string
	Level string
}

// 相关地址信息
type PathConfig struct {
	FilePath   string
	StaticPath string
}

type DbConfig struct {
	Dbtype string
	Sqlite SqliteConfig
	Mysql  mysqlconfig
}
type SqliteConfig struct {
	Filename string
}
type mysqlconfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
}

var c TomlConfig

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("D:\\24123\\code\\go\\L1")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("配置文件不存在!")
		fileName := "config.toml"
		file, _ := os.Create(fileName)
		_, _ = file.WriteString(`
		`)
		defer file.Close()
		fmt.Printf("成功创建 %s 文件,请修改配置文件.", fileName)
		os.Exit(1)
	}

	viper.Unmarshal(&c)
}
func GetConfig() TomlConfig {
	return c
}
