package main

import (
	"encoding/json"
	"log"
)

//change with your nsq topic
const topic = "Roberto_NSQ_Project"

func publish(message Message) (err error) {
	readyMessage, err := json.Marshal(message)
	if err != nil {
		log.Println("[error]Failed to Marshall", err)
		return
	}

	err = NsqProducer.Publish(topic, readyMessage)
	if err != nil {
		log.Println("[error]Failed to publish ", err)
	}
	return
}
