package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/check", check)
	http.ListenAndServe(":8080", nil)
}

func check(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	w.Write([]byte("ok"))
}
