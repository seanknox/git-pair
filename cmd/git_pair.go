package main

import (
	"log"
	"os"

	"github.com/seanknox/git-pair/pkg/userparser"

	"fmt"
)

func main() {
	args := os.Args

	usernames, err := userparser.ParseUsername(args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(usernames)

}
