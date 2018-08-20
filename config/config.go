package config

import (
	"github.com/spf13/viper"
	"fmt"
	"strconv"
)

func GetEnv(key, fallback string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config %s \n", err))
	}
	tmpport := viper.Get("PORT").(int)
	port := strconv.Itoa(tmpport)
	if port == "" {
		return fallback
	}
	return port
}

func GetListenAddress() string {
	port := GetEnv("PORT", "9043")
	return ":" + port
}

func LogSetup() {

}