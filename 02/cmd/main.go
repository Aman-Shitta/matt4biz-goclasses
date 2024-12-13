package main

import (
	"fmt"
	"hello"
	"os"
)

// func main() {
// 	fmt.Printf("Hello, %s!\n", os.Args[1])
// }

// func main() {

// 	if len(os.Args) > 1 {
// 		fmt.Printf("Hello, %s!\n", os.Args[1])
// 	} else {
// 		fmt.Println("Hello, World!")
// 	}
// }

func main() {

	// if len(os.Args) > 1 {
	// 	fmt.Printf(hello.Say(os.Args[1]))
	// } else {
	// 	fmt.Println("Hello, World!")
	// }

	fmt.Println(hello.Say(os.Args[1:]))
}
