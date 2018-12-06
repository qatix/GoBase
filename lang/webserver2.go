package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("this is from string:%v", s))
}

func (s Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("this is from struct:%v", s))
}

func main() {
	// var h Hello
	http.Handle("/", &Hello{})
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

	http.ListenAndServe("localhost:4000", nil)

}
