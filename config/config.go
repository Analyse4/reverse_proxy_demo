package config

import (
	"github.com/spf13/viper"
	"fmt"
	"log"
)

var AConditionUrl string
var BConditionUrl string
var DefaultConditionUrl string


func init()  {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config %s \n", err))
	}
}

func GetEnv(key, fallback string) string {
	port := viper.Get("PORT").(string)
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
	AConditionUrl = viper.Get("REDIRECT_URL.A_CONDITION_URL").(string)
	BConditionUrl = viper.Get("REDIRECT_URL.B_CONDITION_URL").(string)
	DefaultConditionUrl = viper.Get("REDIRECT_URL.DEFAULT_CONDITION_URL").(string)

	//a_condition_url := viper.Get("A_CONDITION_URL")
	//b_condition_url := viper.Get("B_CONDITION_URL")
	//default_condition_url := viper.Get("DEFAULT_CONDITION_URL")

	log.Printf("Server will run on: localhost%s\n", GetListenAddress())
	log.Printf("Redirecting to A url: %s\n", AConditionUrl)
	log.Printf("Redirecting to B url: %s\n", BConditionUrl)
	log.Printf("Redirecting to Default url: %s\n", DefaultConditionUrl)
}