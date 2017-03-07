package user_parser

import (
	"log"
	"os"
)

func ParseUsername() ([]string, error) {
	args := os.Args

	if len(args) < 1 {
		log.Fatal("Please supply at least one GitHub username")
	}

	// username := args[0]
	// return username, nil
	return args, nil
}
