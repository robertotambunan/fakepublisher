package main

import (
	"log"

	nsq "github.com/bitly/go-nsq"
)

var NsqProducer *nsq.Producer

//Change with your nsq address
const nsqAddress = "devel-go.tkpd:4150"

func initialize() error {
	var err error
	config := nsq.NewConfig()
	NsqProducer, err = nsq.NewProducer(nsqAddress, config)
	if err != nil {
		log.Fatal("[error]Failed To Init NSQ", err)
	}
	return err
}
