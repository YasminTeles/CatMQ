// nolint: errcheck
package main

import (
	"log"
	"time"

	"github.com/YasminTeles/CatMQ/badword"
	"github.com/YasminTeles/CatMQ/client"
	"github.com/YasminTeles/CatMQ/message"
)

const (
	Second = 1
	Minute = Second * 60
)

func main() {
	cli := client.NewClientDefault()
	cli.Connect()
	defer cli.Disconnect()

	cli.Consumer()

	badwords, err := badword.NewBadWords()
	if err != nil {
		log.Fatalf("Some error occurred. err: %v", err)
	}

	phrase := cli.Get()

	for isValid(phrase) {
		wait(phrase)

		if !badwords.Check(phrase) {
			cli.Publish(phrase)
		}

		phrase = cli.Get()
	}
}

func isValid(phrase string) bool {
	return phrase != message.MessageError
}

func wait(phrase string) {
	if phrase == "" {
		time.Sleep(Minute * time.Second)
	}
}
