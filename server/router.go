package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)

	for {
		var msg Message
		err := decoder.Decode(msg)
		if err != nil {
			fmt.Println("failed to decode incomming json into msg")
			return
		}
		RouteMessage(msg, conn)

	}

}

func RouteMessage(msg Message, conn net.Conn) {
	switch msg.Type {
	case RegisterType:
		handleRegister(conn, msg)
	case LoginType:
		handleLogin(conn, msg)
	case SendType:
		handleSend(conn, msg)
	case OnlineType:
		handleOnline(conn)
	}

}

func handleRegister(conn net.Conn, msg Message) {
	err := RegisterUser(msg.Username, msg.Password)
	if err != nil {
		sendResponse(conn, ErrorType, err.Error())
		return
	}
	sendResponse(conn, SuccessType, "registeration succesful")
}

func handleLogin(conn net.Conn, msg Message) {
	if err := AuthenticateUser(msg.Username, msg.Password); err != nil {
		sendResponse(conn, ErrorType, err.Error())
		return
	}
	sendResponse(conn, SuccessType, "login successful")
}

func handleSend(conn net.Conn, msg Message) {
	// get user connection object
	resConn, exist := GetUserConnection(msg.Username)
	if !exist {
		sendResponse(conn, ErrorType, "User ofline!")
		return
	}
	// convert msg to json and write into resConn
	json.NewEncoder(resConn).Encode(msg)

}

func handleOnline(conn net.Conn) {
	users := ListUsers()
	resp := Response{
		Type:  OnlineType,
		Users: users,
	}
	// convert resp to json and write it into conn
	json.NewEncoder(conn).Encode(resp)

}

func sendResponse(conn net.Conn, responseType, msg string) {
	resp := Response{
		Type:    responseType,
		Message: msg,
	}
	json.NewEncoder(conn).Encode(resp)
}
