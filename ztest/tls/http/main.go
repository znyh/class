package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "hello, world!\n")

}

func main() {

	http.HandleFunc("/hello", HelloServer)

	err := http.ListenAndServeTLS(":8080", _cert, _key, nil)

	if err != nil {

		log.Fatal("ListenAndServe: ", err)

	}

}

var (
	_cert = ""
	_key  = ""
)
