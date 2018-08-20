package main

import "net/http"

func getEnv(key, fallback string) string {
	return ""
}

func getListenAddress() string {
	port := getEnv("PORT", "9043")
	return ":" + port
}

func LogSetup()  {

}

func requestAndRedirect(w http.ResponseWriter, r *http.Request)  {
	
}

func main()  {
	//log setup values
	LogSetup()

	//start server
	http.HandleFunc("/", requestAndRedirect)
	http.ListenAndServe(getListenAddress(), nil)
}
