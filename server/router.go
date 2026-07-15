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

	for {
		var msg protocol.Message
		err := decoder.Decode(&msg)
		if err != nil {
			fmt.Println("failed to decode incomming json into msg")
			return
		}
		RouteMessage(msg, conn)

	}

}

func RouteMessage(msg protocol.Message, conn net.Conn) {
	switch msg.Type {
	case protocol.RegisterType:
		handleRegister(conn, msg)
	case protocol.LoginType:
		handleLogin(conn, msg)
	case protocol.SendType:
		handleSend(conn, msg)
	case protocol.OnlineType:
		handleOnline(conn)
	}

}

func handleRegister(conn net.Conn, msg protocol.Message) {
	err := RegisterUser(msg.Username, msg.Password)
	if err != nil {
		sendResponse(conn, protocol.ErrorType, err.Error())
		return
	}
	sendResponse(conn, protocol.SuccessType, "registeration succesful")
}

func handleLogin(conn net.Conn, msg protocol.Message) {
	if err := AuthenticateUser(msg.Username, msg.Password); err != nil {
		sendResponse(conn, protocol.ErrorType, err.Error())
		return
	}
	sendResponse(conn, protocol.SuccessType, "login successful")
}

func handleSend(conn net.Conn, msg protocol.Message) {
	// get user connection object
	resConn, exist := GetUserConnection(msg.To)
	if !exist {
		sendResponse(conn, protocol.ErrorType, "User offline!")
		return
	}
	// convert msg to json and write into resConn
	json.NewEncoder(resConn).Encode(msg)

}

func handleOnline(conn net.Conn) {
	users := ListUsers()
	resp := protocol.Message{
		Type:  protocol.OnlineType,
		Users: users,
	}
	// convert resp to json and write it into conn
	json.NewEncoder(conn).Encode(resp)

}

func sendResponse(conn net.Conn, responseType, msg string) {
	resp := protocol.Message{
		Type:    responseType,
		Message: msg,
	}
	json.NewEncoder(conn).Encode(resp)
}
