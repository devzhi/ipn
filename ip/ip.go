package ip

import (
	"io"
	"net/http"
	"os"
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
	// 写入获取到的IP
	f, err := os.OpenFile("old_ip.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	f.Write(body)
	s := string(body)
	return &s, nil
}

/*
读取旧IP方法
*/
func ReadOldIP() (*string, error) {
	f, err := os.Open("old_ip.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fi.Size()
	buffer := make([]byte, filesize)

	bytesread, err := f.Read(buffer)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}

	s := string(bytesread)

	return &s, nil
}
