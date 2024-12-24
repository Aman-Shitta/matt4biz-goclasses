package main

import (
	"fmt"
)

// Pass by refrence (maps [Pointer])
func do(m1 *map[int]int) {
	(*m1)[1] = 3
	*m1 = make(map[int]int) // assigns a new map to m1 (local scope).
	(*m1)[4] = 4
	fmt.Println(m1)
	fmt.Printf("m1@ %p\n", m1)
}

func main() {
	m := map[int]int{4: 1, 7: 2, 8: 3}
	fmt.Printf("m@ %p\n", m)
	do(&m)
	fmt.Printf(" m@ %p\n", m)
	fmt.Println(m)
}

// // Pass by refrence (maps)
// func do(m1 map[int]int) {
// 	m1[1] = 3
// 	m1 = make(map[int]int) // assigns a new map to m1 (local scope).
// 	m1[4] = 4
// 	fmt.Println(m1)
// 	fmt.Printf("m1@ %p\n", m1)
// }

// func main() {
// 	m := map[int]int{4: 1, 7: 2, 8: 3}
// 	fmt.Printf("m@ %p\n", m)
// 	do(m)
// 	fmt.Printf(" m@ %p\n", m)
// 	fmt.Println(m)
// }

// // Pass by refrence (Sllice)
// func do(b []int) int {
// 	b[0] = 0
// 	fmt.Printf("b@ %p\n", b)

// 	return b[1]
// }

// func main() {
// 	a := []int{1, 2, 3}
// 	fmt.Printf("a@ %p\n", a)
// 	v := do(a)

// 	fmt.Println(a, v)
// }

// Pass by value (arrayt)
// func do(b [3]int) int {
// 	b[0] = 0
// 	fmt.Printf("b@ %p\n", b)

// 	return b[1]
// }

// func main() {
// 	a := [3]int{1, 2, 3}
// 	fmt.Printf("a@ %p\n", a)
// 	v := do(a)

// 	fmt.Println(a, v)
// }
