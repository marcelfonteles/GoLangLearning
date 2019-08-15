package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, my name is Inigo Montoya")
}

func inigo(res http.ResponseWriter, req *http.Request) {
	textoBytes := make([]byte, 5)
	file, _ := os.Open("/home/marcelvieira/go/bin/index.html")
	texto, err := file.Read(textoBytes)
	for err == nil {
		fmt.Fprint(res, string(textoBytes[:texto]))
		texto, err = file.Read(textoBytes)
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/inigo", inigo)
	http.ListenAndServe("localhost:3000", nil)
}