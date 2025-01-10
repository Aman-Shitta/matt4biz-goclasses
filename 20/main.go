package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type rupees float32

func (r rupees) String() string {
	return fmt.Sprintf("₹%.2f\n", r)
}

type database map[string]rupees

// handlers

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("Duplicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("Invalid Price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	db[item] = rupees(p)

	fmt.Fprintf(w, "added %s with price %s\n", item, db[item])
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("Entry not found: %q", item)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("Invalid Price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	db[item] = rupees(p)
	fmt.Fprintf(w, "updated %s with new price %s\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("Entry not found: %q", item)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "Deleted item: %s \n", item)
}

func (db database) fetch(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("Entry not found: %q", item)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}

	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}

func main() {

	db := database{
		"shoes": 500,
		"sock":  149,
	}

	// routes
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/fetch", db.fetch)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
