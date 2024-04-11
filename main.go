package main

import (
	"URL_shortener/internal/server"
	"fmt"
	"log"
	"net/http"
)

func main() {

	addr := ":8082"
	srv := server.NewServer(addr)

	fmt.Println("URL Shortener is running...")

	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server error:", err)
	}

	if err := srv.Stop(); err != nil {
		log.Fatal("Server shutdown error:", err)
	}
}
