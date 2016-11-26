package main

import (
  "encoding/json"
  "fmt"
  "github.com/pborman/getopt"
  "net/http"
)


type GithubUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
  getopt.Parse()

  commiter1 := getopt.Arg(0)
  githubGetString := fmt.Sprintf("https://api.github.com/users/%s", commiter1)

	resp, err := http.Get(githubGetString)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	ghu := new(GithubUser)

	json.NewDecoder(resp.Body).Decode(&ghu)

	fmt.Println(ghu)
}
