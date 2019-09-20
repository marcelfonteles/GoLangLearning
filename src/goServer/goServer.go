package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello my name is Inigo Montoya\nAnd you killed my father.")
}

func inigo(res http.ResponseWriter, req *http.Request) {
	file, _ := os.Open("/home/marcelvieira/go/bin/index.html")
	b1 := make([]byte, 510)
	n1, err := file.Read(b1)
	for err == nil {
		fmt.Fprint(res, string(b1[:n1]))
		n1, err = file.Read(b1)
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/inigo", inigo)
	http.ListenAndServe("localhost:8080", nil)
}

