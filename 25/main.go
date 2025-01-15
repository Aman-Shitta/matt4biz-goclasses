package main

import (
	"fmt"
	"time"
)

type T struct {
	i byte
	b bool
}

func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)}

	ch <- t

	t.b = true // UNSAFE AT ANY SPEED
}

func main() {
	vs := make([]T, 5)
	ch := make(chan *T)

	for i := range vs {
		go send(i, ch)
		fmt.Println(i, "sent")
	}

	time.Sleep(1 * time.Second) // all go routines started

	// cop quickyl!
	for i := range vs {
		vs[i] = *<-ch // defrenicng the pointer to get value, so any modifiactions to pointer woin't affct the copy
	}
	// print later
	for _, v := range vs {
		fmt.Println(v)
	}

}
