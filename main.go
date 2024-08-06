package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	db "sales-register/db/sqlc"
	"sales-register/internal/server"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	dbConn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("error connet to db: ", err)
	}

	server := server.NewServer(server.ServerCfg{
		Port:    ":" + port,
		Queries: db.New(dbConn),
	})

	fmt.Println("Server liste on port ", port)
	if err := server.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
