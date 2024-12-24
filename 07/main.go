package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}
		file.Close()
		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
	}

}

// func main() {

// 	for _, fname := range os.Args[1:] {
// 		if file, err := os.Open(fname); err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 		} else {
// 			if data, err := ioutil.ReadAll(file); err != nil {
// 				fmt.Fprintln(os.Stderr, err)
// 			} else {
// 				fmt.Println("The len of file", len(data), "Bytes")
// 			}
// 			file.Close()
// 		}

// 	}
// }

// // FIle I/O
// func main() {
// 	for _, fname := range os.Args[1:] {
// 		file, err := os.Open(fname)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			continue
// 		}

// 		if _, err := io.Copy(os.Stdout, file); err != nil {
// 			fmt.Fprint(os.Stderr, err)
// 			continue
// 		}

// 		file.Close()
// 	}
// }

// Formatting string, int, slices and maps
// func main() {
// 	s := "a string"
// 	b := []byte(s)

// 	// fmt.Printf("%T\n", s)
// 	// fmt.Printf("%q\n", s)
// 	// fmt.Printf("%v\n", s)
// 	// fmt.Printf("%#v\n", s)

// 	fmt.Printf("%T\n", b)
// 	fmt.Printf("%q\n", b)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%#x\n", b)
// 	fmt.Printf("%#v\n", b)
// }

// // func main() {
// // 	m := map[string]int{"and": 1, "or": 2}
// // 	fmt.Printf("%T\n", m)
// // 	fmt.Printf("%v\n", m)
// // 	fmt.Printf("%#v\n", m)
// // }

// // func main() {
// // 	s := []int{1, 2, 3}
// // 	x := [3]rune{'a', 'b', 'c'}

// // 	fmt.Printf("%T\n", s)
// // 	fmt.Printf("%v\n", s)
// // 	fmt.Printf("%#v\n", s)

// // 	fmt.Printf("%T\n", x)
// // 	fmt.Printf("%v\n", x)
// // 	fmt.Printf("%q\n", x)
// // 	fmt.Printf("%#v\n", x)
// // }

// // func main() {
// // 	a, b := 12, 345
// // 	c, d := 1.2, 3.45

// // 	// fmt.Printf("%d %d\n", a, b)
// // 	// fmt.Printf("%#x %#x\n", a, b)
// // 	// fmt.Printf("%#X %#X\n", a, b)

// // 	// fmt.Printf("%f %.3f", c, d)

// // 	fmt.Printf("|%9d|%9d|\n", a, b)
// // 	fmt.Printf("|%-9f|%-9f|\n", c, d)
// // 	fmt.Printf("|%-9d|%-9d|\n", a, b)

// // 	fmt.Printf("|%9f|%9.2f|\n", c, d)

// // }
