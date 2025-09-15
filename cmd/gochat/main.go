package main

import (
	"gochat-frontend/cmd/gochat/handlers"
)

var currentPage string = "users"

func main(){
	serverIP, _ := handlers.GetRCValue("SERVERIP")

	if serverIP == "" {
		serverIP = "gochat.com"
	}

	for{
		handlers.ClearTerminal()

		switch currentPage {
		case "root":
			handlers.RenderRootPage(&currentPage)
		case "create-account":
			handlers.RenderCreateAccountPage(&currentPage, serverIP)
		case "login":
			handlers.RenderLoginPage(&currentPage, serverIP)
		case "users":
			handlers.RenderUsersPage(&currentPage, serverIP)
		case "chat":
			handlers.RenderChatPage(&currentPage, serverIP)

		}
	}
}
