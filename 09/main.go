package main

import (
	"fmt"
)

func fib() func() int {
	a, b := 0, 1 // This is maintained in state for the closure()
	fmt.Println(a, b)

	return func() int {
		// fmt.Printf("Before: a = %d, b = %d\n", a, b)
		a, b = b, a+b
		// fmt.Printf("After: a = %d, b = %d\n", a, b)
		return b
	}
}
func main() {

	f, g := fib(), fib()

	fmt.Println(f(), f(), f(), f(), f(), f())
	fmt.Println(g(), g(), g(), g(), g(), g())

	// f := fib()

	// for x := f(); x < 100; x = f() {
	// 	fmt.Println(x)
	// }

}
