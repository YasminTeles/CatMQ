// nolint: errcheck
package badword

import (
	"log"

	"github.com/YasminTeles/CatMQ/client"
	"github.com/YasminTeles/CatMQ/message"
)

func Start() {
	client.Connect()
	// defer client.Disconnect()

	client.Consumer()

	badwords, err := NewBadWords()
	if err != nil {
		log.Fatalf("Some error occurred. err: %v", err)
	}

	phrase := client.Get()

	for isValid(phrase) {
		if !badwords.Check(phrase) {
			client.Publish(phrase)
		}

		phrase = client.Get()
	}
}

func isValid(phrase string) bool {
	return phrase != "" && phrase != message.MessageError
}
