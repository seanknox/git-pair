package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
