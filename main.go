package main

import "github.com/aditya-gupta-dev/conf/config"

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	var user User
	if err := config.InitConfig(user, "aasdigo", "config.json"); err != nil {
		panic(err)
	}
}
