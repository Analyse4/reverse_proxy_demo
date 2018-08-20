package main

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

func getEnv(key, fallback string) string {
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

func getListenAddress() string {
	port := getEnv("PORT", "9043")
	return ":" + port
}

func LogSetup() {

}

func requestAndRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In func requestAndRedirect")
}

func main() {
	//log setup values
	LogSetup()

	//start server
	http.HandleFunc("/", requestAndRedirect)
	http.ListenAndServe(getListenAddress(), nil)

	//test
	//fmt.Println(getListenAddress())
}
