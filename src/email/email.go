package email

import (
	"log"
	"net/smtp"
)

func Send(body string) {
	from := "tester.123.pepito@gmail.com"
	pass := "adrenalina12"
	to := "losidoh756@whipjoy.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
