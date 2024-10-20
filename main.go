package main

import (
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	// Register the pattern once
	mux.HandleFunc("/users", api.usersHandler)

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
