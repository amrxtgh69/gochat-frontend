package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"io"
	"time"
)

type User struct {
	FullName string
	UserName string
	Password string
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
	fmt.Println(ANSI_GREEN, "Success:", body, ", redirecting to home page in 3 seconds...", ANSI_RESET)
	time.Sleep(3 * time.Second)
}


func RenderUsersPage(currentPage *string) {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println(
		" =========[Users]==========\n",
		"| Create Group[$1]\n",
		"| Logout[$2]\n",
		"==========================",
		)
	fmt.Println(
		" ==========================\n",
		"| bokshi[$3]\n",
		"| amrxtgh[$4]\n",
		"==========================",
		)
	fmt.Println("")
	
	fmt.Print(">>> ")
	reader.Scan()
	if  reader.Text() == "$3" {
		*currentPage = "chat"
		return
	}
}

func RenderChatPage(currentPage *string) {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println(
		" =========[Chat]====================\n",
		"| Go to Users[$1]\n",
		"| Send File[$2]\n",
		"====================================",
		)
	fmt.Println(
		" ====================================\n",
		"| Currently chatting with bokshi\n",
		"====================================\n",
		"| You: Hello\n",
		"| Bokshi: Hi, Bokshi\n",
		"====================================",
		)

	fmt.Println("")
	
	fmt.Print(">>> ")
	reader.Scan()
	if  reader.Text() == "$1" {
		*currentPage = "users"
		return
	}
}
