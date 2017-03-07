package user_parser

import (
	"log"
	"os"
)

// ParseUsername returns an array of user strings or an error if no arguments provided
func ParseUsername() ([]string, error) {
	args := os.Args

	if len(args) < 1 {
		log.Fatal("Please supply at least one GitHub username")
	}

	return args, nil
}
