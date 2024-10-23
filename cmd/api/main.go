package main

import (
	"fmt"
	"log"

	"github.com/ferneediaz/gopher-socials/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	fmt.Println("ADDR from env:", cfg.addr)

	mux := app.mount()
	log.Fatal(app.run(mux))
}
