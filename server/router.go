package main

import (
	"chat-app/protocol"
	"encoding/json"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	client := &Client{
		Conn: conn,
	}

	for {
		var msg protocol.Message

		err := decoder.Decode(&msg)

		if err != nil {
			fmt.Println("failed to decode incomming json into msg, due to client disconnection")
			if client.Username != "" {
				DeleteClient(client.Username)

			}

			return
		}
		RouteMessage(msg, client)

	}

}

func RouteMessage(msg protocol.Message, client *Client) {
	switch msg.Type {
	case protocol.RegisterType:
		handleRegister(client, msg)
	case protocol.LoginType:
		handleLogin(client, msg)
	case protocol.SendType:
		handleSend(client, msg)
	case protocol.OnlineType:
		handleOnline(client)
	case protocol.LogoutType:
		DeleteClient(client.Username)
		client.Conn.Close()

	}

}

func handleRegister(client *Client, msg protocol.Message) {
	err := RegisterUser(msg.Username, msg.Password)
	if err != nil {
		sendResponse(client.Conn, protocol.ErrorType, err.Error())
		return
	}
	sendResponse(client.Conn, protocol.SuccessType, "registeration succesful")
}

func handleLogin(client *Client, msg protocol.Message) {
	if err := AuthenticateUser(msg.Username, msg.Password); err != nil {
		sendResponse(client.Conn, protocol.ErrorType, err.Error())
		return
	}
	client.Username = msg.Username
	AddClient(client)
	sendResponse(client.Conn, protocol.SuccessType, "login successful")
}

func handleSend(sender *Client, msg protocol.Message) {
	// get user connection object
	res, exist := GetUserConnection(msg.To)
	if !exist {
		sendResponse(sender.Conn, protocol.ErrorType, "User offline!")
		return
	}

	// Overwrite the sender's username with their actual username
	msg.From = sender.Username

	// validate client's online status first
	if sender.Username == "" {
		sendResponse(sender.Conn, protocol.ErrorType, "Please login first")
		return
	}
	// convert msg to json and write into resipient Conn
	json.NewEncoder(res.Conn).Encode(msg)

}

func handleOnline(client *Client) {
	users := ListUsers()
	resp := protocol.Message{
		Type:  protocol.OnlineType,
		Users: users,
	}

	// Validate client's online status fiest
	if client.Username == "" {
		sendResponse(client.Conn, protocol.ErrorType, "Please login first")
		return
	}
	// convert resp to json and write it into conn
	json.NewEncoder(client.Conn).Encode(resp)

}

func sendResponse(conn net.Conn, responseType, msg string) {
	resp := protocol.Message{
		Type:   responseType,
		Status: msg,
	}
	json.NewEncoder(conn).Encode(resp)
}
