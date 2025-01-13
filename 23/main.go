package main

import (
	"log"
	"time"
)

func main() {
	log.Println("start")
	const tickrate = 2 * time.Second
	stopper := time.After(5 * tickrate)

	ticker := time.NewTicker(tickrate).C
loop:
	for {
		select {
		case <-ticker:
			log.Println("tick")
		case <-stopper:
			break loop

		}
	}
	log.Println("end")
}
