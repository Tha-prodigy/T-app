package main

import (
	"net"
	"sync"
)

type Client struct {
	Username string
	Conn     net.Conn
}

var (
	// use mutex to prevent concurrent actions from accessing shared resources at a time
	mu sync.RWMutex

	// This is a shared resource that multiple online users can access at the same time via multiple Goroutines
	// and as a result mulptiple read and write actions can be performed on this map simultanously which can cause program to crash due to data race .
	// Hence we use RWMutex which has two lock types; Lock() and RLock() and two unlocks; RUnlock() and Unlock()
	// best practice is to use Lock() for write actions and use RLock() for read actions for optimal performance.
	activeUsers = make(map[string]*Client)
)

// records users that are online. Each user has it's own tcp connection
func AddClient(client *Client) {
	mu.Lock()
	defer mu.Unlock()
	activeUsers[client.Username] = client

}

// removes users that are disconnected and offline
func DeleteClient(userName string) {
	mu.Lock()
	defer mu.Unlock()
	delete(activeUsers, userName)
}

// get users connection line and connection status to know how to reach each user
func GetUserConnection(username string) (*Client, bool) {
	mu.RLock()
	defer mu.RUnlock()
	client, exist := activeUsers[username]
	return client, exist

}

// list online users
func ListUsers() []string {

	mu.RLock()
	defer mu.RUnlock()
	// slice for online users
	users := make([]string, 0, len(activeUsers))
	for username := range activeUsers {
		users = append(users, username)

	}
	return users

}
