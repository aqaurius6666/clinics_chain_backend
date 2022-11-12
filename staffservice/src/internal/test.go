package main

import (
	"fmt"
	"net/smtp"
)

var MAIL_SMTP_HOST = "email-smtp.ap-southeast-1.amazonaws.com"
var MAIL_SMTP_PORT = "587"
var MAIL_SMTP_SENDER = "mt766157@gmail.com"
var MAIL_SMTP_USERNAME = "AKIAV5FTKULOBNIXUONU"
var MAIL_SMTP_PASSWORD = "BOrJobnxU9gBACkSq7SZsexY06LP8lQyUlLJm60pIc30"
var MAIL_SMTP_AWS_REGION = "ap-southeast-1"

func a() {
	msg := []byte("From: mt766157@gmail.com\r\n" +
		"To: 19020366@vnu.edu.vn\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")

	smtpServer := fmt.Sprintf("%s:%s", MAIL_SMTP_HOST, MAIL_SMTP_PORT)
	auth := smtp.PlainAuth("", string(MAIL_SMTP_USERNAME), string(MAIL_SMTP_PASSWORD), string(MAIL_SMTP_HOST))

	err := smtp.SendMail(smtpServer, auth, MAIL_SMTP_SENDER, []string{"19020366@vnu.edu.vn"}, msg)
	if err != nil {
        fmt.Println(err)
    }

    //fmt.Println("Email sent successfully")
}
