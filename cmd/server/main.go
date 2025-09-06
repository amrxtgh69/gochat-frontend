package main

import (
	"fmt"
	"os"
	"os/signal"

	"honnef.co/go/tools/config"
)

func main() {
	fmt.Println("Start SCP chat server...")
	cfg := config.DefaultConfig()
	chatServer := server.NewSCPServer()
}
