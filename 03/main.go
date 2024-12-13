package main

import "fmt"

func main() {
	a := 1
	b := 2.1

	// fmt.Printf("a: %8T %v \n", a, a)
	// fmt.Printf("a: %8T %v \n", b, b)

	fmt.Printf("a: %8T %[1]v \n", a)
	fmt.Printf("a: %8T %[1]v \n", b)

	a = int(b)
	fmt.Printf("a: %8T %[1]v \n", a)

	b = float64(a)
	fmt.Printf("a: %8T %[1]v \n", b)
}
