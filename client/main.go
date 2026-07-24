package main

import (
	"fmt"
	"log"
	"net"
)




func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("failed to connect to :8080")
		return
	}
	
	
	cl := &Client{
	Conn: conn,
	LoggedIn: false,
	Logginstate: make(chan bool),
	}

	


	// use goroutine for this part so it can continue to run on a different thread independently of the main
	go Receiver(cl)
	
    
	for {
		// fmt.Println("MAIN LOOP:", cl.LoggedIn)
		var choice, choice2 string
		var username, password string
		var recipient, text string
		
		// fmt.Println("LoggedIn =", cl.LoggedIn)
		fmt.Println("Main LoggedIn:", cl.LoggedIn)
		// fmt.Printf("MAIN Client address: %p\n", cl)

		if !cl.LoggedIn {
			fmt.Println("================")
			fmt.Println("===== Menu =====")
			fmt.Println()
			fmt.Println("1. Register")
			fmt.Println("2. Login")
			fmt.Println("3. Exit")
			fmt.Println()

			fmt.Println("Please select an option: ")
			fmt.Scanln(&choice)

			switch choice {
			case "1":
				fmt.Println("Choose a Username to get registered: ")
				fmt.Scanln(&username)
				fmt.Println("Create a login password: ")
				fmt.Scanln(&password)
				err := Register(cl, username, password)
				if err != nil {
					log.Println("Error encountered while trying to register user")
					return
				}
			case "2":
				fmt.Println("Enter your Username: ")
				fmt.Scanln(&username)
				fmt.Println("Enter your login password: ")
				fmt.Scanln(&password)
				err := Login(cl, username, password)
				cl.LoggedIn  = <- cl.Logginstate
				if err != nil {
					fmt.Println("Error!, Unable to log user in.")
					return
				}
			case "3":
				fmt.Println("Bye bye...")
				return
			}
			

		} else {
			fmt.Println("========================")
			fmt.Printf("         Welcome %s\n", cl.Username)
			fmt.Println("========================")
			fmt.Println()
			fmt.Println("1. Send Message")
			fmt.Println("2. Request Online users")
			fmt.Println("3. logout")
			fmt.Println("4. Exit")
			fmt.Println()
			fmt.Print("Choice: ")
			fmt.Scanln(&choice2)

			switch choice2 {

			case "1":
				fmt.Println("Who would you like to send a message? ")
				fmt.Scanln(&recipient)
				fmt.Println("Enter Message: ")
				fmt.Scanln(&text)
				err := Chat(cl, recipient, text)
				if err != nil {
					fmt.Println("Error!, Unable to send text")
					return
				}
			case "2":
				err := OnlineUsers(cl)
				if err != nil {
					fmt.Println("unable to get online users")
				}
			case "3":
				err := Logout(cl, cl.Username)
				if err != nil {
					fmt.Println("Unable to logout")
				}
			case "4":
				return

			}

		}

	}
}
