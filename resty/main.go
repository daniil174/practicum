package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	client := resty.New()

	var users []User
	url := "https://jsonplaceholder.typicode.com/users"

	_, err := client.R().SetResult(&users).Get(url)

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%#v\n", users)

	var out []string
	for _, v := range users {
		out = append(out, v.Username)
	}
	fmt.Println(strings.Join(out, ` `))
}
