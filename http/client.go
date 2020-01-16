package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://127.0.0.1:8080/hello"

	curl1(url)
	curl2(url)
	curl3(url)
}

func curl1(url string) {
	//resp, err := http.Get("http://example.com/")
	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	//resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("response code: %d, body: %s\n", resp.StatusCode, body)
}

func curl2(url string) {
	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("response code: %d, body: %s\n", resp.StatusCode, body)
}

func curl3(url string) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}, // 可不指定
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("response code: %d, body: %s\n", resp.StatusCode, body)
}
