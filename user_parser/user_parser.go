package user_parser

import "fmt"

// ParseUsername returns an array of user strings or an error if no arguments provided
func ParseUsername(usernames []string) ([]string, error) {
	if len(usernames) < 1 {
		return []string{""}, fmt.Errorf("Please supply at least one GitHub username")
	}

	return usernames, nil
}
