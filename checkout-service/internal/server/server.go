package server

import (
	"fmt"
	"log"
	"net/http"
)

func Run() error {
	mux := http.NewServeMux()
	// Checkout Service
	mux.HandleFunc("/checkout/placeorder", checkoutPlaceOrderHandler)
	port := 8081
	log.Printf("Server is running on port %d...\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
