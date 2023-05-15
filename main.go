package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to the database
	db := ConnectDB()
	defer db.Close()

	// Set up the router
	router := mux.NewRouter()

	// Register the handlers
	router.HandleFunc("/register", RegisterHandler(db)).Methods("POST")
	router.HandleFunc("/login", LoginHandler(db)).Methods("POST")
	router.HandleFunc("/get-all-users", GetAllUsersHandler(db))

	// Set up CORS headers
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", cors(router))

	// Start the server
	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
