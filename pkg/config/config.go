package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func NewConfig() *viper.Viper {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		flag.StringVar(&envConf, "conf", "config/local.yml", "config path, eg: -conf config/local.yml")
		flag.Parse()
	}
	if envConf == "" {
		envConf = "local"
	}
	fmt.Println("load conf file:", envConf)
	return GetConfig(envConf)

}
func GetConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	conf.SetDefault("gorm-gen.dal", "./internal/dal")
	conf.SetDefault("gorm-gen.model", "./internal/model")
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
