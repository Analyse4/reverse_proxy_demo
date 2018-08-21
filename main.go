package main

import (
	"net/http"
	//"github.com/Analyse4/reverse_proxy_demo/config"
	"reverse_proxy_demo/config"
	"reverse_proxy_demo/handler"
)

func main() {
	//log setup values
	config.LogSetup()
	//start server
	http.HandleFunc("/", handler.RequestAndRedirect)
	http.ListenAndServe(config.GetListenAddress(), nil)
}
