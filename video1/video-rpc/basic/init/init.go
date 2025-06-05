package init

import (
	"github.com/spf13/viper"
	"log"
	"video1/video-rpc/basic/config"
	"video1/video-rpc/utils"
)

func init() {
	InitConfig()
	InitDB()

}

func InitDB() {
	config.DB = utils.GlobalMysql()
}

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		return
	}
	if err := v.Unmarshal(&config.GlobalMysql); err != nil {
		return
	}

	log.Println("viper init success")
}
