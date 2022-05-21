package gu

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func CommonPostJson(url string, headers map[string]string, body map[string]interface{}) ([]byte, error) {
	payloadByte, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	payload := bytes.NewReader(payloadByte)

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	for s, a := range headers {
		req.Header.Add(s, a)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	return resBody, err
}
func CommonGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	return resBody, err
}
