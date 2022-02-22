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

func HttpPost(url string, obj interface{}) error {
	_, err := http.Post(url, "application/json", postBody(obj))
	if err != nil {
		return err
	}
	return nil
}

func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bd, nil
}

func HttpPut(url string, obj interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, postBody(&obj))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func HttpDelete(url string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return err
}

func CalculaImc(peso float64, altura float64) float64 {
	if peso == 0 || altura == 0 {
		return 0
	}
	return peso / (math.Pow(altura, 2))
}