package main

import (
	"fmt"
	"sort"
)

// Sorting
type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

func (s Organs) Len() int {
	return len(s)
}

func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByWeight struct {
	Organs
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

type ByName struct {
	Organs
}

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func main() {
	organs := Organs{{"brain", 1340}, {"liver", 1494}, {"spleen", 162}, {"pancreas", 131}, {"heart", 290}}

	fmt.Println(organs)

	sort.Sort(ByWeight{organs})
	fmt.Println(organs)

	sort.Sort(ByName{organs})
	fmt.Println(organs)
}

// -----------------------------------------------
// type Pair struct {
// 	Path string
// 	Hash string
// }

// func (p Pair) String() string {
// 	return fmt.Sprintf("Hash of %s is %s\n", p.Path, p.Hash)
// }

// type PairWithLength struct {
// 	Pair
// 	Length int
// }

// func (p PairWithLength) String() string { // is not overriding, as it is not inherticance but composition
// 	return fmt.Sprintf("Hash of %s is %s with length of %d\n ", p.Path, p.Hash, p.Length)
// }

// func (p Pair) Filename() string { // method
// 	return filepath.Base(p.Path)
// }

// func Filename(p Pair) string { // function
// 	return filepath.Base(p.Path)
// }

// type FileNamer interface {
// 	Filename() string
// }

// func main() {

// 	p := Pair{"/usr", "0xdeafbeef"}
// 	fmt.Println(p)

// 	// pl := PairWithLength{Pair{""}, 121}
// 	pl := PairWithLength{Pair{"/usr/lib", "0xdeadbeef"}, 121}
// 	fmt.Println(pl)

// 	// fmt.Println(Filename(p))
// 	// fmt.Println(Filename(pl)) // ERROR : (type PairWithLength has no field or method Filename)
// 	// fmt.Println(Filename(pl.Pair))

// 	fmt.Println(p.Filename())
// 	fmt.Println(pl.Filename())

// 	var fn FileNamer = PairWithLength{Pair{"/usr/lib", "0xdeadbeef"}, 121}
// 	fmt.Println(fn.Filename())
// }
