package utils

import (
	"bytes"
	"net/http"
)

type HttpUtil struct {
}

var httpUtil = new(HttpUtil)

func (h *HttpUtil) PostRequest(url string, jsonData []byte) *http.Response {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	return BaseRequest(request)
}

func (h *HttpUtil) GetRequest(url string) *http.Response {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	return BaseRequest(request)
}

func BaseRequest(request *http.Request) *http.Response {
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	return response
}
