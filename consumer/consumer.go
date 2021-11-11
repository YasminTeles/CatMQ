// nolint: errcheck
package main

import (
	"log"

	"github.com/YasminTeles/CatMQ/badword"
	"github.com/YasminTeles/CatMQ/client"
	"github.com/YasminTeles/CatMQ/message"
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
		if !badwords.Check(phrase) {
			cli.Publish(phrase)
		}

		phrase = cli.Get()
	}
}

func isValid(phrase string) bool {
	return phrase != "" && phrase != message.MessageError
}
