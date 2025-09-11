package main

import (
	"github.com/amritxtgh69/gochat-frontend/cmd/gochat/handlers"
)

var currentPage string = "root"

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
		}
	}
}
