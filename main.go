package main

import (
	"fmt"
	"log"
	"net/http"

	"mcserver/handlers"
)


func main() {
	http.HandleFunc("/api/server", handlers.ServerInfoHandler)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
