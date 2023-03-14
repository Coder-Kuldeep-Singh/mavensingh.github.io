package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var cache sync.Map

func main() {
	http.HandleFunc("/", handler)
	log.Println(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	v, ok := cache.Load(r.URL.Path)
	if ok {
		fmt.Fprintln(w, v)
		return
	}

	response := time.Now().UTC().Format(time.RFC3339)
	cache.Store(r.URL.Path, response)
	fmt.Fprintln(w, response)
}
