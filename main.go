package main

import (
	"fmt"
	"net/http"
	//"github.com/Analyse4/reverse_proxy_demo/config"
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
