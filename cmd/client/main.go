package main

import (
	"fmt"
	"os"
	"scp-chat/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: Client <username>")
		os.Exit(1)
	}

	username := os.Args[1]
	cfg := config.DefaultConfig()
	fmt.Printf("Connecting to server at %s...\n", cfg.GetServerAddress())
}
