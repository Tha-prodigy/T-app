package main

import (
	// "bufio"
	"chat-app/protocol"
	"fmt"
	"log"
)

// func readLine(prompt string) string {
// 	reader := bufio.NewReader(os.Stdin)
// 	text, _ := reader.ReadString('\n')
// 	str := strings.TrimSpace(text)
// 	return str

// }

func Register(client *Client, username, password string) error {
	msg := protocol.Message{
		Type:     protocol.RegisterType,
		Username: username,
		Password: password,
	}
	err := Send(client, msg)
	if err != nil {
		log.Println("unable to send msg to connection")
		return err
	}
	return nil

}

func Login(client *Client, username, password string) error {
	msg := protocol.Message{
		Type:     protocol.LoginType,
		Username: username,
		Password: password,
	}
	return Send(client, msg)
}

func Chat(client *Client, receiver, text string) error {
	msg := protocol.Message{
		Type: protocol.MessageType,
		To:   receiver,
		Body: text,
	}
	err := Send(client, msg)
	if err != nil {
		log.Println("Unable to send chart over the connection")
		return err
	}
	return nil
}

func OnlineUsers(cl *Client) error {
	msg := protocol.Message{
		Type: protocol.OnlineType,
	}
	err := Send(cl, msg)
	if err != nil {
		log.Println("Unable to retreive online users")
		return err
	}
	return nil
}

func Logout(client *Client, username string) error {
	msg := protocol.Message{
		Type:     protocol.LogoutType,
		Username: username,
	}
	err := Send(client, msg)
	if err != nil {
		fmt.Println("Error while login out")
		return err
	}
	client.LoggedIn = false
	return nil

}
