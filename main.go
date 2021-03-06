package main

import (
	"net/http"
	//"github.com/Analyse4/reverse_proxy_demo/config"
	"reverse_proxy_demo/config"
	"reverse_proxy_demo/handler"
)

func main() {
	// Log setup values
	config.LogSetup()
	// Start server
	http.HandleFunc("/", handler.RequestAndRedirect)
	http.ListenAndServe(config.GetListenAddress(), nil)
}
