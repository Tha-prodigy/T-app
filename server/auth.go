package main

import "fmt"

type User struct {
	Username string
	Password string
}

var users = map[string]User{}

func RegisterUser(username, password string) error {
	_, ok := users[username]
	if ok {
		return fmt.Errorf("Nigga that name already in the system")
	}
	users[username] = User{Username: username,
		Password: password}
	return nil

}

func AuthenticateUser(username, password string) error {
	u, ok := users[username]
	if !ok {
		return fmt.Errorf("Nigga gtfo you don't belong here!")

	}
	if u.Password != password {
		return fmt.Errorf("Invalid password!")
	}
	return nil

}
