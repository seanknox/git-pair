package main

import "net/http"
import "encoding/json"
import "fmt"

type GithubUser struct {
  Name string `json:"name"`
  Email string `json:"email"`
}

func main() {
	resp, err := http.Get("https://api.github.com/users/ryanswood")

  if err != nil {
    panic(err.Error())
  }

	defer resp.Body.Close()

  ghu := new(GithubUser)

  json.NewDecoder(resp.Body).Decode(&ghu)

  fmt.Println(ghu)
}

  // body, err := ioutil.ReadAll(resp.Body)

  // if err != nil {
  //   panic(err.Error())
  // }

  // fmt.Print(body)
