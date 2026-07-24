package main

import "net"

type Client struct {
    Conn net.Conn

    Username string

    LoggedIn bool
    Logginstate chan bool
}