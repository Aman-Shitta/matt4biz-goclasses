package main

import "fmt"

type errFoo struct {
	err  error
	path string
}

func (f errFoo) Error() string {
	return fmt.Sprintf("%s : %s", f.path, f.err)
}

func XYZ(a int) *errFoo { // returns a concrete pointer to an errFoo
	return nil
}

func main() {

	var err error = XYZ(1) // BAD: interface gets a nil concrete ptr

	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("OK!")
	}

}
