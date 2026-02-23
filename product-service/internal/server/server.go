package server

import (
	"fmt"
	"log"
	"net/http"
)

func Run() error {
	mux := http.NewServeMux()
	// Product Service
	mux.HandleFunc("/products", productListHandler)
	mux.HandleFunc("/products/{id}", productDetailHandler)
	port := 8082
	log.Printf("Server is running on port %d...\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
