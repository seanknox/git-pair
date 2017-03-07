package user_parser

import (
	"fmt"
	"os"
)

// ParseUsername returns an array of user strings or an error if no arguments provided
func ParseUsername() ([]string, error) {
	args := os.Args

	if len(args) < 1 {
		return []string{""}, fmt.Errorf("Please supply at least one GitHub username")
	}

	return args, nil
}
