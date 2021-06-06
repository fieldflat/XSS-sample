package main

import (
	"fmt"
	"net/http"
)

func attackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Scheme, r.URL.Path, r.URL.Query().Get("sid"))
}

func main() {
	http.HandleFunc("/", attackHandler)
	http.ListenAndServe(":8081", nil)
}
