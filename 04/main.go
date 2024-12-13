package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	s := "elite"
// 	fmt.Printf("%8T %[1]v, %v\n", s, len(s)) // string elite, 5

// 	fmt.Printf("%8T %[1]v\n", []rune(s)) // []int32 [101 108 105 116 101]

// 	b := []byte(s)
// 	fmt.Printf("%8T %[1]v %v\n", b, len(b)) // []uint8 [101 108 105 116 101] 5

// 	c := "你好!"
// 	fmt.Printf("%8T %[1]v %v\n", c, len(c)) // string 你好! 7
// 	d := []rune(c)

// 	fmt.Printf("%8T %[1]v %v\n", d, len(d)) // []int32 [20320 22909 33] 3

// 	b = []byte(c)
// 	fmt.Printf("%8T %[1]v %v\n", b, len(b)) // []uint8 [228 189 160 229 165 189 33] 7
// }

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not engough arguments")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}

}
