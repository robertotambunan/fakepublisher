package main

import (
	"log"
)

func main() {
	log.Println("Initializing...")
	initialize()
	log.Println("Publishing...")
	err := handlerPublish()
	if err != nil {
		log.Println("[error][warnig]Failed to do anything")
	}
}

func handlerPublish() (err error) {
	message := createPreMessage("random")
	for _, msg := range message {
		err = publish(msg)
		if err != nil {
			log.Println("[error][main]Failed to publish")
		}
	}
	return
}
