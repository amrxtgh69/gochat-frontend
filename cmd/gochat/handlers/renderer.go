package handlers

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	fullName string
	userName string
	password string
}

func RenderRootPage(currentPage *string) {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("==========[ROOT]==========")
	fmt.Println("")
	fmt.Println("Login[$1]")
	fmt.Println("Create Account[$2]")

	fmt.Print("> ")
	reader.Scan()

	if reader.Text() == "$1" {
		*currentPage = "login"
		return
	} else if reader.Text() == "$2"{
		*currentPage = "create-account"
		return
	} else {
		fmt.Print("Invalid option")
	}
}

func RenderCreateAccountPage(currentPage *string)  {
	reader := bufio.NewScanner(os.Stdin)
	var user User

	fmt.Println("==========[CREATE ACCOUNT]==========")
	
	fmt.Println("")
	fmt.Println("NAVIGATE -----▶ ROOT[$1]")
	fmt.Println("")
	
	fmt.Println("Full Name")
	fmt.Print("> ")
	reader.Scan()
	if  reader.Text() == "$1" {
		*currentPage = "root"
		return
	}
	user.fullName = reader.Text()

	fmt.Println("User Name")
	fmt.Print("> ")
	reader.Scan()
	if reader.Text() == "$1" {
		*currentPage = "root"
		return
	}
	user.userName = reader.Text()

	fmt.Println("Password")
	fmt.Print("> ")
	reader.Scan()
	if reader.Text() == "$1" {
		*currentPage = "root"
		return
	}
	user.password = reader.Text()
	
	// TODO: use net/http to send user credentials to server
}

func RenderLoginPage(currentPage *string) {
	reader := bufio.NewScanner(os.Stdin)
	var user User

	fmt.Println("==========[LOGIN]==========")
	
	fmt.Println("")
	fmt.Println("NAVIGATE -----▶ ROOT[$1]")
	fmt.Println("")
	
	fmt.Println("User Name")
	fmt.Print("> ")
	reader.Scan()
	if  reader.Text() == "$1" {
		*currentPage = "root"
		return
	}
	user.userName = reader.Text()

	fmt.Println("Password")
	fmt.Print("> ")
	reader.Scan()
	if reader.Text() == "$1" {
		*currentPage = "root"
		return
	}
	user.password = reader.Text()

	// TODO: use net/http to send user credentials to server
}
