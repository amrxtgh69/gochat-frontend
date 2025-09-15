package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"io"
	"time"
	"encoding/json"
)

type User struct {
	FullName string
	UserName string
	Password string
}

type Message struct {
    ID       int    `json:"id"`
    Sender   string `json:"sender"`
    Receiver string `json:"receiver"`
    Content  string `json:"content"`
}

var ANSI_GREEN string = "\033[1;30;42m" // bold black text with green bg
var ANSI_RED string = "\033[1;41m" // bold text with red bg
var ANSI_RESET string = "\033[0m"

func RenderRootPage(currentPage *string) {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println(
		" ==========[ROOT]===========\n",
		"| Go to Login[$1]\n",
		"| Go to Create Account[$2]\n",
		"===========================",
		)
	fmt.Println("")
	
	fmt.Print(">>> ")
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

func RenderCreateAccountPage(currentPage *string) {
	reader := bufio.NewScanner(os.Stdin)
	var user User

	fmt.Println(
		" =====[Create Account]=====\n",
		"| Go to Root[$1]\n",
		"==========================",
	)
	fmt.Println("")

	fmt.Println("Full Name")
	fmt.Print(">>> ")
	reader.Scan()
	if reader.Text() == "$1" {
		*currentPage = "root"
		return
	}
	user.FullName = reader.Text()

	fmt.Println("User Name")
	fmt.Print(">>> ")
	reader.Scan()
	if reader.Text() == "$1" {
		*currentPage = "root"
		return
	}
	user.UserName = reader.Text()

	fmt.Println("Password")
	fmt.Print(">>> ")
	reader.Scan()
	if reader.Text() == "$1" {
		*currentPage = "root"
		return
	}
	user.Password = reader.Text()

	data := fmt.Sprintf(`{"fullName":"%s","userName":"%s","password":"%s"}`, user.FullName, user.UserName, user.Password)
	resp, err := http.Post("http://127.0.0.1:8080/create-account", "application/json", strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := strings.TrimSpace(string(bodyBytes))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println(ANSI_RED, "Err:", body, ", refreshing in 3 seconds...", ANSI_RESET)
		time.Sleep(3 * time.Second)
		return
	}

	*currentPage = "login"
	fmt.Println(ANSI_GREEN, "Success:", body, ", redirecting to login in 3 seconds...", ANSI_RESET)
	time.Sleep(3 * time.Second)
}

func RenderLoginPage(currentPage *string) {
	reader := bufio.NewScanner(os.Stdin)
	var user User

	fmt.Println(
		" =========[Login]==========\n",
		"| Go to Root[$1]\n",
		"| Forgot Password[$2]\n",
		"==========================",
	)
	fmt.Println("")

	fmt.Println("User Name")
	fmt.Print(">>> ")
	reader.Scan()
	if reader.Text() == "$1" {
		*currentPage = "root"
		return
	}
	user.UserName = reader.Text()

	fmt.Println("Password")
	fmt.Print(">>> ")
	reader.Scan()
	if reader.Text() == "$1" {
		*currentPage = "root"
		return
	}
	user.Password = reader.Text()

	data := fmt.Sprintf(`{"userName":"%s","password":"%s"}`, user.UserName, user.Password)
	resp, err := http.Post("http://127.0.0.1:8080/login", "application/json", strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := strings.TrimSpace(string(bodyBytes))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println(ANSI_RED, "Err:", body, ", refreshing in 3 seconds...", ANSI_RESET)
		time.Sleep(3 * time.Second)
		return
	}

	*currentPage = "users"
	setRCValue("USERNAME", user.UserName) // save username to .gochatrc
	fmt.Println(ANSI_GREEN, "Success:", body, ", redirecting to home page in 3 seconds...", ANSI_RESET)
	time.Sleep(3 * time.Second)
}

func RenderUsersPage(currentPage *string) {
	username, _ := getRCValue("USERNAME")
	if username == "" {
		*currentPage = "root" // no saved user
		return
	} else {
		*currentPage = "users" // auto-login
	}

    resp, err := http.Get("http://127.0.0.1:8080/users")
    if err != nil {
        fmt.Println("Error fetching users:", err)
        time.Sleep(2 * time.Second)
        return
    }
    defer resp.Body.Close()

    bodyBytes, _ := io.ReadAll(resp.Body)
    var users []User
    if err := json.Unmarshal(bodyBytes, &users); err != nil {
        fmt.Println("Error parsing users:", err)
        return
    }

    reader := bufio.NewScanner(os.Stdin)

    fmt.Println(" =========[Users]==========")
    fmt.Println("| Create Group[$1]")
    fmt.Println("| Logout[$2]")
    fmt.Println("==========================")

    // Dynamically list users with numbers $3, $4, $5 ...
    for i, user := range users {
        fmt.Printf("| %s[$%d]\n", user.UserName, i+3)
    }
    fmt.Println("==========================")
    fmt.Println("")

    fmt.Print(">>> ")
    reader.Scan()
    input := reader.Text()

    if input == "$1" {
        // create group handler
		return
    } else if input == "$2" {
        *currentPage = "root"
        return
    } else {
        // map input to correct user index
        for i, user := range users {
            if input == fmt.Sprintf("$%d", i+3) {
				*currentPage = "chat"
				fmt.Println("Now chatting with", user.UserName)
				setRCValue("RECEIVERUSERNAME", user.UserName) // save receiver
				return
            }
        }
    }
}

func RenderChatPage(currentPage *string) {
	reader := bufio.NewScanner(os.Stdin)
	
	currentUser, _ := getRCValue("USERNAME")
	receiver, _ := getRCValue("RECEIVERUSERNAME")
	if receiver == "" {
		fmt.Println("No receiver selected!")
		*currentPage = "users"
		return
	}

	for {
		// Fetch messages from backend
		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/get-messages?sender=%s&receiver=%s", currentUser, receiver))
		if err != nil {
			fmt.Println("Error fetching messages:", err)
			time.Sleep(2 * time.Second)
			continue
		}
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var msgs []Message
		json.Unmarshal(bodyBytes, &msgs)

		// Render chat
		fmt.Println("=========[Chat]=========")
		fmt.Println("| Go to Users[$1]")
		fmt.Println("========================")
		for _, m := range msgs {
			if m.Sender == currentUser {
				fmt.Printf("[You -> %s]: %s\n", receiver, m.Content)
			} else {
				fmt.Printf("[%s -> You]: %s\n", m.Sender, m.Content)
			}
		}
		fmt.Println("========================")

		// Input new message
		fmt.Print(">>> ")
		reader.Scan()
		text := reader.Text()

		if text == "$1" {
			*currentPage = "users"
			return
		} else if text == "" {
			continue
		}

		// Send to backend
		msg := Message{Sender: currentUser, Receiver: receiver, Content: text}
		data, _ := json.Marshal(msg)
		resp2, err := http.Post("http://127.0.0.1:8080/send-message", "application/json", strings.NewReader(string(data)))
		if err != nil {
			fmt.Println("Error sending message:", err)
			time.Sleep(2 * time.Second)
		} else {
			resp2.Body.Close()
		}
	}
}
