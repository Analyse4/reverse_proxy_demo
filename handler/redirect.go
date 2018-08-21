package handler

import (
	"net/http"
	"encoding/json"
	"reverse_proxy_demo/models"
	"io/ioutil"
	"log"
	"bytes"
	"strings"
	"reverse_proxy_demo/config"
	"net/url"
	"net/http/httputil"
)

func parseRequestBody(req *http.Request) models.RequestPayloadStruct {
	decoder := requestBodyDecoder(req)

	var requestpayload models.RequestPayloadStruct
	err := decoder.Decode(&requestpayload)
	if err != nil{
		panic(err)
	}
	return requestpayload
}

// Get a json decoder for given requests body
func requestBodyDecoder(req *http.Request) *json.Decoder {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Err reading body: %v\n", err)
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

// Given a request send it to appropriate url
func RequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	// Parse request for payload
	payload := parseRequestBody(r)
	redirecturl := getProxyUrl(payload.ProxyCondition)
	log.Printf("proxy_condition: %s, proxy_url: %s\n", payload.ProxyCondition, redirecturl)
	serverReverseProxy(redirecturl, w, r)
}

// Get proxy-url from req-payload
func getProxyUrl(rp string) string {
	proxycondition := strings.ToUpper(rp)

	if proxycondition == "A" {
		return config.AConditionUrl
	} else if proxycondition == "B" {
		return  config.BConditionUrl
	}else {
		return config.DefaultConditionUrl
	}
}

// Serve a reverse-proxy for a given url
func serverReverseProxy(target string, res http.ResponseWriter, req *http.Request)  {
	// Parse the url
	url, _ := url.Parse(target)
	// Create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}