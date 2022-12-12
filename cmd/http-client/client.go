package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "echo",
	}

	b, err := u.MarshalBinary()
	if err != nil {
		return
	}

	target := string(b)

	get(target)
	post(target)
}

func get(u string) {
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}

func post(u string) {
	b := []byte("Here is some payload")

	resp, err := http.Post(u, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}
