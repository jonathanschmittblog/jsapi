package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

func CheckError(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}

func postBody(obj interface{}) *bytes.Buffer {
	postBody, err := json.Marshal(&obj)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(postBody)
}

func HttpPost(url string, obj interface{}) int {
	println("url:" + url)
	resp, err := http.Post(url, "application/json", postBody(obj))
	if err != nil {
		println(err.Error())
		return 0
	}
	return resp.StatusCode
}

func HttpGet(url string) (int, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, nil
	}
	bd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil
	}
	return resp.StatusCode, bd
}

func HttpPut(url string, obj interface{}) int {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, postBody(&obj))
	if err != nil {
		return 0
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	return resp.StatusCode
}

func HttpDelete(url string) int {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return 0
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	return resp.StatusCode
}

func CalculaImc(peso float64, altura float64) float64 {
	if peso == 0 || altura == 0 {
		return 0
	}
	return peso / (math.Pow(altura, 2))
}