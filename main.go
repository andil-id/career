package main

import (
	"log"

	"github.com/segmentio/ksuid"
)

func main() {
	id := ksuid.New().String()
	log.Println(id)
}
