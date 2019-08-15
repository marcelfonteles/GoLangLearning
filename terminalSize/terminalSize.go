package main

import (
	"os/exec"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	fmt.Printf("out: %#v\n", string(out))
	fmt.Printf("err: %#v\n", err)
	if err != nil {
		log.Fatal(err)
	}

	indexHeight := strings.Index(string(out), " ")
	indexWidth := strings.Index(string(out), "\n")

	heightString := string(out[0:indexHeight])
	widthString  := string(out[indexHeight+1:indexWidth])


}