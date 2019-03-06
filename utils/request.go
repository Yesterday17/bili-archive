package utils

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const referer = "https://www.bilibili.com"

func Request(method, url, cookies string, reqBody io.Reader, header map[string]string) (*http.Response, error) {
	transport := &http.Transport{
		DisableCompression:  true,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: transport,
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Referer", referer)
	req.Header.Set("Cookie", cookies)

	if header != nil {
		for _, item := range header {
			req.Header.Set(item, item)
		}
	}

	return client.Do(req)
}

func Get(url, cookies string, reqBody io.Reader) (string, error) {
	res, err := Request("GET", url, cookies, reqBody, nil)
	if err != nil {
		return "", err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetJson(url, cookies string, v interface{}) error {
	if content, err := Get(url, cookies, nil); err != nil {
		return err
	} else {
		return json.Unmarshal([]byte(content), &v)
	}
}
