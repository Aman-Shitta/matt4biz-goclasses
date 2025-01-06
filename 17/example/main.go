package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(l)
	return l, nil
}

func main() {
	var dst ByteCounter

	f1, _ := os.Open("a.txt")
	// f2, _ := os.Create("out.txt")
	c := &dst

	n, _ := io.Copy(c, f1)
	// n, _ := io.Copy(f2, f1)
	// Copy takes in Writer and Reader Type interface
	// so we can create our own Writer compatible type

	fmt.Println("copied", n, "bytes")
	fmt.Println(dst)
}
