package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "You called foo!")
}

func bar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "You called bar!")
}

func baz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "You called baz!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", foo)
	r.HandleFunc("/bar", bar)
	r.NotFoundHandler = http.HandlerFunc(baz)
	http.ListenAndServe(":3333", r)
}
