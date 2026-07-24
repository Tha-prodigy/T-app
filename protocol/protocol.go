package protocol

const (
	RegisterType = "register"
	LoginType    = "login"
	LogoutType 	 = "logout"
	SendType     = "send"
	OnlineType   = "online"
	SuccessType	 = "success"
	ErrorType    = "error"
	MessageType  = "message"
)

type Message struct {
	Type     string `json:"type,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	To       string `json:"to,omitempty"`
	From 	 string	`json:"from,omitempty"`
	Body     string `json:"body,omitempty"`
	Users   []string `json:"users,omitempty"`
	Status  string	 `json:"status,omitempty"`
}
