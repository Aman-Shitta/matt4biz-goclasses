package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

const URL = "https://jsonplaceholder.typicode.com/"

type todo struct {
	UserID    int    `json:"userID"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
	<h1> Todo #{{.Id}}</h1>
	<div>{{ printf "User %d" .UserID}}</div>
	<div>{{ printf "%s (completed : %t)" .Title .Completed }}</div>
`

func handler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(URL + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// var item todo
	// err = json.Unmarshal(body, &item)

	var item todo
	if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.New("mine")

	tmpl.Parse(form)

	tmpl.Execute(w, item)

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

// SERVER -----------------------
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World! from %s", r.URL.Path[1:])
// }

// func main() {

// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":8000", nil))

// }

// CLIENT -----------------------

// func main() {

// 	resp, err := http.Get(URL + "todos/1/")

// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(-1)
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		body, err := io.ReadAll(resp.Body)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			os.Exit(-1)
// 		}

// 		var r todo
// 		err = json.Unmarshal(body, &r)
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			os.Exit(-1)
// 		}

// 		fmt.Printf("%#v\n", r)
// 	}
// }
