package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mcserver/handlers"
)

func main() {
	
	http.HandleFunc("/api/server", handlers.ServerInfoHandler) // http://localhost:8080/api/server?ip=host:port
	http.HandleFunc("/api/list", handlers.ServerListHandler)// http://localhost:8080/api/list?id=1
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
