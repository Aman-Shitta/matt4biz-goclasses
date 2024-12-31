package main

import (
	"fmt"
	// "time"
	"encoding/json"
)

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
	// words []string `json:"words,omitempty"` // go vet warning (words cannot be exported because first char in field name is lower)
}

// type Response struct {
// 	Page  int      // `json:"page"`
// 	Words []string // `json:"words,omitempty"`
// } // {"Page":1,"Words":null}

func main() {

	// r1 := Response{Page: 1, Words: []string{"in", "up"}}
	// {"page":1,"words":["in","up"]}

	// r1 := Response{Page: 1, Words: []string{}}
	// {"page":1}

	r1 := Response{Page: 1}
	// {"page":1}

	j, _ := json.Marshal(r1)
	fmt.Println(string(j))
	fmt.Printf("%#v\n", r1)

	var r2 Response

	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)

}

// type Employee struct {
// 	Name   string
// 	Number int
// 	Boss   *Employee
// 	Hired  time.Time
// }

// func main() {
// 	var album = struct {
// 		title  string
// 		artist string
// 		year   int
// 		copies int
// 	}{
// 		"The white album",
// 		"The Beatles",
// 		1968,
// 		10000000000,
// 	}

// 	fmt.Println(album)

// }

// func main() {

// 	c := map[string]*Employee{}

// 	c["Payal"] = &Employee{"Payal", 1, nil, time.Now()}
// 	c["Aman"] = &Employee{"Aman", 2, c["Payal"], time.Now()}

// 	fmt.Printf("%T %+[1]v\n", c["Payal"])
// 	fmt.Printf("%T %+[1]v\n", c["Aman"])
// 	// fmt.Printf("%T %#[1]v\n", emp)
// }

// func main() {

// 	emp := Employee{}
// 	emp.Name = "Aman"
// 	emp.Number = 1
// 	emp.Hired = time.Now()

// 	// boss := Employee{"Payal", 2, nil, time.Now()}
// 	// emp.Boss = &boss

// 	boss := &Employee{"Payal", 2, nil, time.Now()}
// 	emp.Boss = boss

// 	fmt.Printf("%T %+[1]v\n", emp)
// 	// fmt.Printf("%T %#[1]v\n", emp)
// }
