# Slices in detail

Slice -> diagram (https://youtu.be/pHl9r3B2DFI?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&t=332)

func main() {

	var s []int

	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
	//outputs: 0, 0, []int,  true, []int(nil)

	t := []int{}
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil)
	//outputs: 0, 0, []int, false, []int{}

	u := make([]int, 5)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil)
	// outputs: 5, 5, []int, false, []int{0, 0, 0, 0, 0}

	v := make([]int, 0, 5)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil)
	// outputs: 0, 5, []int, false, []int{}

}

Slice descreiptor

              len    cap   addr
                ^     ^     ^
                |     |     |
 slice --->  [..4..|..4..|..0xab..]   ------> [0|1|2|3]
               
               SLICE DESCRIPTIOR



# Empty vs Nil slice

- Slices (and maps) encoding differently in JSON when nil

import (
    "fmt"
    "encoding/json"
)
func main() {
    var a []int

    j1, _ := json.Marshal(a)
    fmt.Println(string(j1)) // null

    b := []int{}

    j2, _ := json.Marshall(b)
    fmt.Println(string(j2)) // []
}

*When checking if an array is empty or not

// WROONG
if a := nil {
    ...
    // only pass if nil array, and not for empty array
}

// Right
if len(a) == 0 {
    ...
    // Checks out for both nil and empty arrays.
}

**You can append to null slice


# Len vs Capacity


func main() {
    a := [3]int{1, 2, 3}

    b := a[0:1]

    fmt.Println("a :", a)
    fmt.Println("b :", b)
}


func main() {
	a := [3]int{1, 2, 3}

	b := a[0:1]

	fmt.Println("a :", a)
	fmt.Println("b :", b)

	c := b[0:2] // WTF

	fmt.Println("c :", c)

	fmt.Printf("a == %d %[1]T %d %d @%p\n", a, len(a), cap(a), &a)
	fmt.Printf("b == %d %[1]T %d %d @%p\n", b, len(b), cap(b), &b)
	fmt.Printf("c == %d %[1]T %d %d @%p\n", c, len(c), cap(c), &c)
}

"When c is assigned to b[0:2], it still points to the first two elements of a, so any modifications to c[n], where n <= 1, will modify the underlying a[n] as well."

    - Slices (b and c) share the same underlying array (a), so modifying c will directly affect the corresponding elements of a.

"The two colon syntax will limit the capacity of the new slice, i.e., b := a[0:1:1] (arr[i:j:k] means {len = j-i} and {cap = k-i}). The restriction will not allow c to get references of the underlying array, hence c := b[0:2] will go out of bounds and panic."

    - The two-colon syntax (a[i:j:k]) restricts the slice's capacity to k-i. If b := a[0:1:1] is created with cap(b) = 1, any attempt to create c := b[0:2] will exceed b's capacity and result in a runtime panic.


https://youtu.be/pHl9r3B2DFI?list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6