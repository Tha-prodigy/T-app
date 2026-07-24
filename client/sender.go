package main

import (
	"chat-app/protocol"
	"encoding/json"
	// "net"
)

func Send(client *Client, msg protocol.Message) error {
	encoder := json.NewEncoder(client.Conn)
	err := encoder.Encode(msg)
	if err != nil {
		return err

	}
	return nil

}

