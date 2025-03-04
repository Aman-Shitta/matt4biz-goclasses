package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	for i := 0; i < 12; i++ {
		// i1 := <-chans[0]
		// i2 := <-chans[1]

		select {
		case m0 := <-chans[0]:
			log.Println("recieved :: 0 :: ", m0)
		case m1 := <-chans[1]:
			log.Println("recieved :: 1 :: ", m1)
			// default:
			// 	log.Println("default")
		}
	}
}
