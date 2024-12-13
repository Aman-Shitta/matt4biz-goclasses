package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello, %s\n!", os.Args[1])
}


-> os.Args[0] would be the name of the program
-> os.Args[n] would be the command line argumnets after program, where n > 0 


Test cases for go

- a function for test always starts with "Test" followed by any name.
- test function should always have the argument of pointer type to "testing.T"
- In case of failed scenario use "Errof" to show failed cases with custom message.

