package g

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type ContentType string

const (
	ContentTypeJSON ContentType = "json"
)

func Post(reqUrl string, data map[string]string, contentType ContentType) (resp *http.Response, err error) {
	var params []byte
	headers := map[string]string{}
	if contentType == ContentTypeJSON {
		params, err = json.Marshal(data)
		if err != nil {
			return
		}
		headers["content-type"] = "application/json;charset=utf-8"
	} else {
		vs := url.Values{}
		for k, v := range data {
			vs.Add(k, v)
		}
		params = []byte(vs.Encode())
	}

	return Request("POST", reqUrl, bytes.NewBuffer(params), headers)
}

func Get(url string) (resp *http.Response, err error) {
	return Request("GET", url, nil)
}

func Request(method, url string, body io.Reader, headers ...map[string]string) (resp *http.Response, err error) {
	var req *http.Request
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	if len(headers) > 0 {
		hs := headers[0]
		for k, v := range hs {
			req.Header.Add(k, v)
		}
	}
	return Client.Do(req)
}
