package main

import (
	"log"

	"github.com/ferneediaz/gopher-socials/internal/db"
	"github.com/ferneediaz/gopher-socials/internal/env"
	"github.com/ferneediaz/gopher-socials/internal/store"
)

const version = "0.0.1"
//	@title	Gopher Socials API

//	@description	API FOR gopher social.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/v1
//
//@securityDefinitions.api_key 	ApiKeyAuth
//@in 				header
//@name 			Authorization
//@description 	    Authorization
func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "http://localhost:8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Printf("database connection pool established")
	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
