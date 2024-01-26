package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var conf *viper.Viper

func init() {
	conf = NewConfig()
}
func GetString(key string) string {
	return conf.GetString(key)
}
func GetStringSlice(key string) []string {
	return conf.GetStringSlice(key)
}
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
	conf.SetDefault("gorm-gen.model", "./model")
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
