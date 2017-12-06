package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func readFile(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}

	str := string(data)

	return strings.Split(str, "\n")
}
