package ip

import (
	"io"
	"net/http"
)

/*
获取IP信息方法
*/
func GetIPInfo() (*string, error) {
	resp, err := http.Get("http://myip.ipip.net/")
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	s := string(body)
	return &s, nil
}
