package config

import (
	"github.com/spf13/viper"
	"log"
)

type MysqlConfig struct {
	Hostname    string
	Port        string
	Username    string
	Database    string
	Password    string
	Charset     string
	ParseTime   string
	Loc         string
	MaxIdleConn int
	MaxOpenConn int
	AxLifetime  int
}

/**
当查到使用  user  会返回列表
没有查到表示空值
*/
func (mysqlConfig *MysqlConfig) GetMysqlConfig() (config MysqlConfig) {
	viper.SetConfigName("db")
	viper.SetConfigType("ini")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	config.Hostname = viper.GetString("mysql.hostname")
	config.Password = viper.GetString("mysql.password")
	config.Port = viper.GetString("mysql.port")
	config.Database = viper.GetString("mysql.database")
	config.Username = viper.GetString("mysql.username")
	config.Charset = viper.GetString("mysql.charset")
	config.ParseTime = viper.GetString("mysql.parse_time")
	config.Loc = viper.GetString("mysql.loc")
	config.MaxIdleConn = viper.GetInt("mysql.max_idle_conn")
	config.MaxOpenConn = viper.GetInt("mysql.max_open_conn")
	config.AxLifetime = viper.GetInt("mysql.ax_lifetime")
	return config
}
