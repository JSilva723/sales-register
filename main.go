package main

import (
	"fmt"
	"log"
	"sales-register/internal/server"
)

func main() {
	server := server.NewServer(server.ServerCfg{
		Port: ":3040",
	})

	fmt.Println("Server liste on port 3040")
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
