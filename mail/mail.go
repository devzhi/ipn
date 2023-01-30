package mail

import (
	"ipn/config"
	"log"
	"net/smtp"
	"strconv"
)

func SendIP(ip string) {

	config := config.GetConfig()

	to_arr := []string{config.To}

	message := []byte("From: IPN <" + config.From + ">\r\n" +
		"To: " + config.To + "\r\n" +
		"Subject: 【IPN】设备 " + config.Device + " 检测到IP地址变更\r\n" +
		"\r\n" + ip +
		"\r\n")

	auth := smtp.PlainAuth("", config.From, config.Password, config.SmtpHost)

	err := smtp.SendMail(config.SmtpHost+":"+strconv.Itoa(config.SmtpPort), auth, config.From, to_arr, message)
	if err != nil {
		log.Println("邮件发送失败", err)
	} else {
		log.Println("邮件发送成功")
	}
}
