package main

import (
	"log"
	"os"

	"fmt"

	. "github.com/seanknox/git-pair/user_parser"
)

func main() {
	args := os.Args

	usernames, err := ParseUsername(args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(usernames)

}
