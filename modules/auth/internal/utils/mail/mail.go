package mail

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(to, subject, body string) error {

	mailer := gomail.NewMessage()
	mailer.SetHeader("To", to)
	mailer.SetHeader("From", os.Getenv("SENDER_MAIL"))
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)
	dialer := gomail.NewDialer(os.Getenv("MAIL_HOST"), 587, os.Getenv("SENDER_MAIL"), os.Getenv("SENDER_PASSWORD"))
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}
	return nil
}
