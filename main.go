package main

import (
	"ipn/ip"
	"ipn/mail"
	"log"
	"os"
)

func main() {
	// 获取旧IP
	old_ip, err := ip.ReadOldIP()
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}
	// 获取当前IP
	curr_ip, err := ip.GetIPInfo()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(*curr_ip)
	if old_ip == nil || *curr_ip != *old_ip {
		log.Println("发现IP变更")
		mail.SendIP(*curr_ip)
	} else {
		log.Println("未发现IP变更")
	}
}
