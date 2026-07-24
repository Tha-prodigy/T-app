package main

import (
	"chat-app/protocol"
	"encoding/json"
	"fmt"
	// "net"
)

func Receiver(client *Client) {
	decoder := json.NewDecoder(client.Conn)
	for {
		var msg protocol.Message
		if err := decoder.Decode(&msg); err != nil {
			fmt.Println("Disconnected from server")

			return
		}
		// fmt.Printf("DEBUG: %#v\n", msg)
		switch msg.Type {
		case protocol.OnlineType:
			fmt.Println("Online Users")
			fmt.Println("-------------")
			for _, user := range msg.Users {
				fmt.Println("-", user)

			}
		case protocol.SuccessType:
			fmt.Println(msg.Status)
			if msg.Status == "login successful" {
				// fmt.Println("DEBUG: Switching LoggedIn to true")
				fmt.Printf("RECEIVER Client address: %p\n", client)
				client.LoggedIn = true
				client.Username = msg.Username
				client.Logginstate <- true

			}
			fmt.Println("Receiver LoggedIn:", client.LoggedIn)
		case protocol.ErrorType:
			fmt.Println("Error!", msg.Status)
		case protocol.MessageType:
			fmt.Printf("\n[%s]: %s\n", msg.From, msg.Body)
		default:
			fmt.Println("Unknown message type", msg.Type)

		}
	}
}
