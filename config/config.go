package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf *viper.Viper

func InitConfig() (err error) {
	Conf = viper.New()
	Conf.SetConfigName("config")    // 文件名
	Conf.SetConfigType("yaml")      // 文件类型
	Conf.AddConfigPath("../config") // 搜索路径

	err = Conf.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	return
}
