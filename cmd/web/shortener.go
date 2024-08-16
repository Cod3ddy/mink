package main

import (
	"math/rand"
	"net/http"
	"time"
)

func shorten(url string) (string, error){
	code, err := pingURL(url)
	if err != nil{
		return "", err
	}

	return http.StatusText(code), nil
}

func generateShortKey() string {
	const charset = "acbdefghijklmnopqrstuvwxyz0123456789"
	const keyLength = 6

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	shortKey := make([]byte, keyLength)
	for i := range shortKey{
		shortKey[i] = charset[r.Intn(len(charset))]
	}
	return string (shortKey)
}

func pingURL(url string)(int, error){
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil{
		return 0, err
	}
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return 0, err
	}

	resp.Body.Close()

	return resp.StatusCode, nil
}