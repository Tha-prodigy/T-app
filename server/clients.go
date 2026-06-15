package main

import (
	"net"
	"sync"
)

var (
	// use mutex to prevent concurrent actions from accessing shared resources at a time
	mu sync.RWMutex

	activeUsers = make(map[string]net.Conn)
)

// records users that are online. Each user has it's own tcp connection
func AddClient(username string, conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()
	activeUsers[username] = conn

}

// removes users that are disconnected and offline
func DeleteClient(userName string) {
	mu.Lock()
	defer mu.Unlock()
	delete(activeUsers, userName)
}

// get users connection line and connection status to know how to reach each user
func GetUserConnection(username string) (net.Conn, bool) {
	mu.RLock()
	defer mu.RUnlock()
	conn, exist := activeUsers[username]
	return conn, exist

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
