package server

import (
	"fmt"
	"log"
	"net/http"
)

func Run() error {
	mux := http.NewServeMux()
	// Auth Service
	mux.HandleFunc("/auth/login", authLoginHandler)
	mux.HandleFunc("/auth/logout", authLogoutHandler)
	port := 8080
	log.Printf("Server is running on port %d...\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
