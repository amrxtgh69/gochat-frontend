package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

type User struct {
	Fullname string 
	Username string
	Password string 
}

var users = make(map[string]User)
var CurrentUser *User

func main() {
	users["admin"] = User{Fullname: "Administrator", Username: "admin", Password: "foobar123"}
	for {
		fmt.Println("===gochat===")
		fmt.Println("1. Login")
		fmt.Println("2. Create Account")
		fmt.Println("3. Exit")
		fmt.Println("Choose Option")
		
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			login()
		case 2:
			createAccount()
		case 3:
			fmt.Println("GOODBYE!!")
			os.Exit(0)
		default:
			fmt.Println("Invalid option!!")
		}
	}
}

func login() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Password: ")
	bytepassword, _ := term.ReadPassword(int(syscall.Stdin))
	password := string(bytepassword)
	fmt.Println()

	user, exists := users[username]
	if !exists {
		fmt.Println("User not found!!")
		return
	}
	if user.Password != password {
		fmt.Println("Invalid password")
		return
	}
	CurrentUser = &user
	fmt.Printf("Welcome %s!\n", user.Fullname)
	chatPrompt()
}


func createAccount() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("FullName: ")
	scanner.Scan()
	fullname := scanner.Text()

	fmt.Print("Username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Password: ")
	bytepassword, _ := term.ReadPassword(int(syscall.Stdin))
	password := string(bytepassword)

	if _, exists := users[username]; exists {
		fmt.Println("Username already exists!!")
		return
	}
	newUser := User{
		Fullname: fullname,
		Username: username,
		Password: password,
	}
	users[username] = newUser
	
	fmt.Println("Account created successfully!!")
	fmt.Printf("There are %d users in gochat", len(users))
}

func chatPrompt() {
	return
}
