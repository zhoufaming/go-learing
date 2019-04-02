package main

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello World"))

	})

	http.ListenAndServe(":8888", nil)

}


// package main
// import (
// "fmt"
// "net/http"
// )
// func main() {
// http.HandleFunc("/", hello)
// http.ListenAndServe("localhost:8000", nil)
// }
// func hello(w http.ResponseWriter, r *http.Request) {
// fmt.Fprintln(w, "Hello, Gophers! Welcome to Qcon!")
// }
