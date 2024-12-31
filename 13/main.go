package main

import "fmt"

func main() {

	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}, {6, 7}}

	a := [][]byte{}

	for _, item := range items {
		a = append(a, item[:])
	}

	fmt.Printf("items : %d, with type %[1]T\n", items)
	fmt.Printf("a : %d, witn type %[1]T\n", a)
}
