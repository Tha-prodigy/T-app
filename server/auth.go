package main

import "fmt"

type User struct {
	Username string
	Password string
}

var users = map[string]User{}

// this registers new users
func RegisterUser(username, password string) error {
	_, ok := users[username]
	if ok {
		fmt.Println("this username already exists")
		return fmt.Errorf("Nigga that name already in the system")
	}
	users[username] = User{Username: username,
		Password: password}
	return nil

}

// authenticates users name and password while logging in 
func AuthenticateUser(username, password string) error {
	u, ok := users[username]
	if !ok {
		fmt.Println("invalid username")
		return fmt.Errorf("Nigga gtfo you don't belong here!")

	}
	if u.Password != password {
		fmt.Println("invalid password")
		return fmt.Errorf("Invalid password! mf")
	}
	return nil

}
