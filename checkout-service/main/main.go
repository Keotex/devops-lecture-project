package main

import (
	"fmt"
	"os"

	"github.com/Keotex/devops-lecture-project/checkout-service/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
