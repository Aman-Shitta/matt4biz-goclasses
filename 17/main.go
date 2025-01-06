package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

func (l *Line) ScaleBy(f float64) {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
}

func main() {
	side := Line{Point{2, 5}, Point{6, 8}}
	fmt.Println(side.Distance())
	side.ScaleBy(3)
	fmt.Println(side.Distance())

	// fmt.Println(&Line{Point{2, 5}, Point{6, 8}}.ScaleBy(2).Distance())                // wring because first SclaeBy won't return anything to run Distance on and The value is Literal type and not pointer

}

//  -------------------------------------------
// type Point struct {
// 	X, Y float64
// }

// type Line struct {
// 	Begin, End Point
// }

// type Path []Point

// func (l Line) Distance() float64 {
// 	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
// }

// func (p Path) Distance() (sum float64) {
// 	for i := 1; i < len(p); i++ {
// 		sum += Line{p[i-1], p[i]}.Distance()
// 	}
// 	return
// }

// type Distancer interface {
// 	Distance() float64
// }

// func PrintDistance(d Distancer) {
// 	fmt.Println(d.Distance())
// }

// func main() {
// 	side := Line{Point{2, 5}, Point{6, 8}}
// 	// fmt.Println(side.Distance())
// 	PrintDistance(side)

// 	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}

// 	// fmt.Println(perimeter.Distance())
// 	PrintDistance(perimeter)
// }

// --------------------------------------

// type IntSlice []int // named type

// func (is IntSlice) String() string {
// 	var strs []string

// 	for _, v := range is {
// 		strs = append(strs, strconv.Itoa(v))
// 	}

// 	return "[" + strings.Join(strs, ";") + "]"
// }

// func main() {

// 	var v IntSlice = []int{1, 2, 3}
// 	var s fmt.Stringer = v

// 	for i, x := range v {
// 		fmt.Printf("%d: %d\n", i, x)
// 	}
// 	fmt.Printf("%T %[1]v\n", v)
// 	// prints main.IntSlice [1;2;3]
// 	fmt.Printf("%T %[1]v\n", s)
// 	// prints main.IntSlice [1;2;3]
// }
