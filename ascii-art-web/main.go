package main

import (
	src "ascii-art-web/src"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", src.RootHandler)
	mux.HandleFunc("/ascii-art", src.AsciiHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server starting on port :8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	fmt.Println("Server closed.")
}
