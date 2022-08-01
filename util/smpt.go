package util

import (
	"log"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

var (
	smtpDialer   *gomail.Dialer
	emailAddress string
)

func InitUtilSMTP() {
	emailAddress = os.Getenv("SMTP_EMAIL")
	emailPassword := os.Getenv("SMTP_PASSWORD")
	smtpAddress := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")

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
