package main

import (
	"fmt"
	"net/http"
	"reverse_proxy_demo/config"
)



func requestAndRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In func requestAndRedirect")
}

func main() {
	//log setup values
	config.LogSetup()

	//start server
	http.HandleFunc("/", requestAndRedirect)
	http.ListenAndServe(config.GetListenAddress(), nil)
}
