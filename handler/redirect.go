package handler

import (
	"net/http"
	"encoding/json"
	"reverse_proxy_demo/models"
	"io/ioutil"
	"log"
	"bytes"
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

//get a json decoder for given requests body
func requestBodyDecoder(req *http.Request) *json.Decoder {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Err reading body: %v\n", err)
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

//given a request send it to appropriate url
func RequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	//parse request for payload
	payload := parseRequestBody(r)
	log.Println(payload.ProxyCondition)
}