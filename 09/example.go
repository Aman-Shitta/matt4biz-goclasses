package main

import "fmt"

func do(f func()) {
	f()
}

func main() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}

//  For anyone using Go >=1.22, the behavior where the variable i
// captures the same location and prints the same value (e.g., 4)
// when a closure is called has been fixed in the latest versions.
