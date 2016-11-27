package main

import (
	"encoding/json"
	"fmt"
	"github.com/pborman/getopt"
	"net/http"
	"os"
)

type GithubUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	getopt.Parse()

	if getopt.NArgs() == 0 {
		fmt.Println("must specify a git author")
		os.Exit(1)
	}

	author := getopt.Arg(0)
	githubHttpGetString := fmt.Sprintf("https://api.github.com/users/%s", author)

	resp, err := http.Get(githubHttpGetString)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	githubAuthorDetails := new(GithubUser)

	json.NewDecoder(resp.Body).Decode(&githubAuthorDetails)

	fmt.Println(githubAuthorDetails)
}
