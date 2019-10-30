package config

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Admin Admin
}

// Admin ...
type Admin struct {
	Username string
	Password string
	Path     string
	DBName   string
	Config   string
}

// DBConfig ...
var DBConfig Config

func init() {
	v := viper.New()
	v.SetConfigName("config")             // 设置配置文件名(不带后缀)
	v.AddConfigPath("./config/dbconfig/") // 第一个搜索路径
	v.SetConfigType("json")               // 配置文件格式
	err := v.ReadInConfig()               // 搜索路径,并读取配置数据
	if err != nil {
		log.Printf("Fatal error config file: %+v\n", err)
		os.Exit(-1)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	if err := v.Unmarshal(&DBConfig); err != nil {
		log.Println("Unmarshal config file err:", err)
		os.Exit(-1)
	}
}
