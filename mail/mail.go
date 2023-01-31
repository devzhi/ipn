package mail

import (
	"ipn/config"
	"log"
	"net/smtp"
	"strconv"
	"strings"
)

func SendIP(ip string) {

	config := config.GetConfig()

	// 构建收件人格式
	var recives strings.Builder
	for i := 0; i < len(config.To); i++ {
		if i != 0 {
			recives.WriteString(",")
		}
		recives.WriteString(config.To[i])
	}

	// 构建邮件格式(RFC822)
	message := []byte("From: IPN <" + config.From + ">\r\n" +
		"To: " + recives.String() + "\r\n" +
		"Subject: 【IPN】设备 " + config.Device + " 检测到IP地址变更\r\n" +
		"\r\n" + ip +
		"\r\n")

	// 认证
	auth := smtp.PlainAuth("", config.From, config.Password, config.SmtpHost)

	// 发送邮件
	err := smtp.SendMail(config.SmtpHost+":"+strconv.Itoa(config.SmtpPort), auth, config.From, config.To, message)
	if err != nil {
		log.Println("邮件发送失败", err)
	} else {
		log.Println("邮件发送成功")
	}
}
