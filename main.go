package main

import (
	"fmt"
	"log"
	"sudoku-solver/server"
)

func main() {
	port := "8080"
	err := server.StartServer(port)
	if err == nil {
		fmt.Printf("Server is running on port %s\n", port)
	}
	log.Fatalf("Failed to start server: %v", err)
}
