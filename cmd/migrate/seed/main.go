package main

import (
	"fmt"
	"log"

	"github.com/ferneediaz/gopher-socials/internal/db"
	"github.com/ferneediaz/gopher-socials/internal/env"
	"github.com/ferneediaz/gopher-socials/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable")
	fmt.Println("Database Address:", addr) // Add this line
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
