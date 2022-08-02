package util

import (
	"log"
	"strconv"

	"github.com/go-gomail/gomail"
	"github.com/wikankun/user-service/config"
)

var (
	smtpDialer   *gomail.Dialer
	emailAddress string
)

func InitUtilSMTP() {
	emailAddress = config.Config.SMTP.Email
	emailPassword := config.Config.SMTP.Password
	smtpAddress := config.Config.SMTP.Host
	smtpPortStr := config.Config.SMTP.Port

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Println("util: cannot parse smtp port " + smtpPortStr)
		log.Panic(err)
	}

	smtpDialer = gomail.NewDialer(smtpAddress, smtpPort, emailAddress, emailPassword)
}

func SendEmail(receiver string, subject string, content string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", emailAddress)
	mail.SetHeader("To", receiver)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", content)
	return smtpDialer.DialAndSend(mail)
}
