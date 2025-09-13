package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
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

	data := fmt.Sprintf("{\"fullName\": \"%s\", \"userName\": \"%s\", \"password\": \"%s\"}", user.fullName, user.userName, user.password)
	resp, err := http.Post("http://127.0.0.1/create-account", "application/json", strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 400 {
		println("Err:", resp.Body, ", refressing the page in 3 second.")
		time.Sleep(3 * time.Second)
		return;
	}

	*currentPage = "login"
	println("Success:", resp.Body, ", redirecting to the login page in 3 second.")
	time.Sleep(3 * time.Second)
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
