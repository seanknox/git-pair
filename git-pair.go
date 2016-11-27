package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	// client initializes a default http client for connecting to the github api
	client *http.Client

	// Timeout specifies the time (in seconds) before an http request should be
	// canceled
	timeout = 10 * time.Second
)

// GithubUser is the name and email of a github user
type GithubUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func init() {
	client = &http.Client{
		Timeout: timeout,
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("expected git author as first argument to 'git-pair'")
	}

	username := args[1]
	user, err := fetchUserDetails(username)
	if err != nil {
		log.Fatalf("error fetching github details for %s: %s", username, err)
	}

	fmt.Printf("User Details\nName: \t%s\nEmail: \t%s\n", user.Name, user.Email)
}

// fetchUserDetails retrieves the user contact information from the github api.
// Requests taking longer than 10 seconds will be canceled.
func fetchUserDetails(username string) (user *GithubUser, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s", username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
