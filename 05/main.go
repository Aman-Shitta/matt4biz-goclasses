package main

import (
	"os"
	"fmt"
	"sort"
	"bufio"
)
// func main() {
// 	t := []byte("string")

// 	fmt.Println(len(t), t)
// 	fmt.Println(t[2])
// 	fmt.Println(t[:2])               // start at 0, upto 2(exclude 2, i.e, [0, 1])
// 	fmt.Println(t[3:5], len(t[3:5])) // start at 3, upto 5(exclude 5, i.e, [3, 4])
// }

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")


	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{key: k, val: v})

	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val || len(ss[i].key) > len(ss[j].key)
	})

	for _, v := range ss {
		fmt.Println(v.key, v.val)
	}
}
