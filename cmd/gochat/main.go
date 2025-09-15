package main

import (
	"gochat-frontend/cmd/gochat/handlers"
)

var currentPage string = "users"

func main(){
	for{
		handlers.ClearTerminal()

		switch currentPage {
		case "root":
			handlers.RenderRootPage(&currentPage)
		case "create-account":
			handlers.RenderCreateAccountPage(&currentPage)
		case "login":
			handlers.RenderLoginPage(&currentPage)
		case "users":
			handlers.RenderUsersPage(&currentPage)
		case "chat":
			handlers.RenderChatPage(&currentPage)

		}
	}
}
