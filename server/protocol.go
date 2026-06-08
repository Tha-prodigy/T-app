package main

const (
	RegisterType = "register"
	LoginType    = "login"
	SendType     = "send"
	OnlineType   = "online"
)

type Message struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	To       string `json:"to"`
	Body     string `json:"body"`
}
