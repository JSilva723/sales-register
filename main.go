package main

import (
	"fmt"
	"log"
	"os"
	"sales-register/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	server := server.NewServer(server.ServerCfg{
		Port: ":" + port,
	})

	fmt.Println("Server liste on port ", port)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
