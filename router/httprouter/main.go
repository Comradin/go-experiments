package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index main index
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("Index called")
	fmt.Fprint(w, "Welcome!\n")
}

// Hello echo reponse
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Hello called")
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func NotFound(w http.ResponseWriter, r *http.Request) httprouter.Handle {
	http.Redirect(w, r, "/", 302)
}

func main() {
	router := httprouter.New()
	router.NotFound = NotFound

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
