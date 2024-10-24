package main

import (
	"log"

	"github.com/ferneediaz/gopher-socials/internal/env"
	"github.com/ferneediaz/gopher-socials/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewStorage(nil)
	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
