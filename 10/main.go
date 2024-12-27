package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}

	b := a[0:1]

	fmt.Println("a :", a)
	fmt.Println("b :", b)

	c := b[0:3] // WTF
	c[0] = 9
	fmt.Printf("a[%p] %v\n", &a, a)
	fmt.Printf("b[%p] %[1]v\n", b)
	fmt.Printf("c[%p] %[1]v\n", c)
	c = append(c, 5)

	fmt.Printf("\na[%p] %v\n", &a, a)
	fmt.Printf("b[%p] %[1]v\n", b)
	fmt.Printf("c[%p] %[1]v\n", c)
}

// func main() {
// 	a := [3]int{1, 2, 3}

// 	b := a[0:1]

// 	fmt.Println("a :", a)
// 	fmt.Println("b :", b)

// 	c := b[0:3] // WTF

// 	fmt.Printf("a[%p] %v\n", &a, a)
// 	fmt.Printf("b[%p] %[1]v\n", b)
// 	fmt.Printf("c[%p] %[1]v\n", c)
// 	c = append(c, 5)

// 	fmt.Printf("\na[%p] %v\n", &a, a)
// 	fmt.Printf("b[%p] %[1]v\n", b)
// 	fmt.Printf("c[%p] %[1]v\n", c)
// }

// func main() {
// 	a := [3]int{1, 2, 3}

// 	b := a[0:1]

// 	fmt.Println("a :", a)
// 	fmt.Println("b :", b)

// 	c := b[0:3] // WTF

// 	fmt.Println("c :", c)

// 	fmt.Printf("a == %d %[1]T %d %d @%p\n", a, len(a), cap(a), &a)
// 	fmt.Printf("b == %d %[1]T %d %d @%p\n", b, len(b), cap(b), &b)
// 	fmt.Printf("c == %d %[1]T %d %d @%p\n", c, len(c), cap(c), &c)
// }

// func main() {

// 	var s []int

// 	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
// 	//outputs: 0, 0, []int,  true, []int(nil)

// 	t := []int{}
// 	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil)
// 	//outputs: 0, 0, []int, false, []int{}

// 	u := make([]int, 5)
// 	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil)
// 	// outputs: 5, 5, []int, false, []int{0, 0, 0, 0, 0}

// 	v := make([]int, 0, 5)
// 	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil)
// 	// outputs: 0, 5, []int, false, []int{}

// }
