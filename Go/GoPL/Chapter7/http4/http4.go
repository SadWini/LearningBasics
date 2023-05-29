package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type database map[string]dollars

type dollars float32

var mux sync.Mutex

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.createUpdate)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) createUpdate(w http.ResponseWriter, req *http.Request) {
	// syntax: /create?item=cpu&price=5
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if item != "" && price != "" {
		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "item: %q, price: %q\n", item, price)
			return
		}
		// create item
		mux.Lock()
		db[item] = dollars(p)
		mux.Unlock()
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprintf(w, "item: %q, price: %q\n", item, price)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	itemTable.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		mux.Lock()
		delete(db, item)
		mux.Unlock()
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
