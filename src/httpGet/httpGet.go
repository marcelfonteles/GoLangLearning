package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Site: ")
	endereco, _ := reader.ReadString('\n')
	endereco = strings.TrimSuffix(endereco, "\n")

	resp, _ := http.Get(endereco)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	file, _ := os.Create("/home/marcelvieira/site.html")
	defer file.Close()
	file.Write(body)
	resp.Body.Close()
}