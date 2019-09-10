package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// data := url.Values{"start": {"0"}, "offset": {"xxx"}}
// headers map[string]interface{}; headers := map[string]string{"name":"pj", "age":24}
func HTTPDo(method, urlpath string, data []byte, headers map[string]string) string {
	method = strings.ToUpper(method)
	mds := []string{"POST", "GET"}
	var flag bool = false
	for _, v := range mds {
		if v == method {
			flag = true
			break
		}
	}
	if !flag {
		return ""
	}

	var req *http.Request
	var err error

	if string(data) != "" {
		if method == "POST" {
			body := bytes.NewBuffer(data)
			req, err = http.NewRequest(method, urlpath, body)
		}
		if method == "GET" {
			body := strings.NewReader(string(data))
			req, err = http.NewRequest(method, urlpath, body)
		}
	} else {
		req, err = http.NewRequest(method, urlpath, nil)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// TODO 待处理
	req.Header.Set("Pl", "02")

	clt := http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	respBody := string(content)
	return respBody
}
