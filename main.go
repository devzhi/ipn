package main

import (
	"ipn/config"
	"ipn/ip"
	"ipn/mail"
	"log"
	"os"

	"github.com/robfig/cron/v3"
)

func main() {
	if config.GetConfig().EnableCron {
		c := cron.New()
		_, err := c.AddFunc(config.GetConfig().CronExp, check)
		if err != nil {
			log.Fatalln(err)
		}
		c.Start()
		log.Println("定时任务启动成功")
		var block chan struct{} //nil channel
		<-block
	} else {
		check()
	}
}

func check() {
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
	} else if config.GetConfig().AlwaysSend {
		log.Println("未发现IP变更，强制发送邮件")
		mail.SendIP(*curr_ip)
	} else {
		log.Println("未发现IP变更")
	}
}
