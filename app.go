package main

import (
	"log"
)

func main() {
	log.Println("Initializing...")
	err := initialize()
	if err != nil {
		log.Fatal("Failed Initialize!", err)
	}
	log.Println("Publishing...")
	err = handlerPublish()
	if err != nil {
		log.Println("[error][warning]Failed to do anything")
	}
}

func handlerPublish() (err error) {
	message := createPreMessage("random")
	for _, msg := range message {
		err = publish(msg)
		if err != nil {
			log.Println("[error][main]Failed to publish", err)
		}
	}
	return
}
