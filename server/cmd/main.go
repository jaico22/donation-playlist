package main

import (
	"encoding/json"
	"fmt"
	"internal/users"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Donation Playlist running")
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/api/users/{facebookUserId}", userController)

	fmt.Println("Listening on port 5000")
	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}

func userController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	requestVars := mux.Vars(r)
	facebookUserId := requestVars["facebookUserId"]
	if r.Method == http.MethodPost {
		newUserRequest := users.NewUserRequest{FacebookUserId: facebookUserId}
		responce := users.NewUser(newUserRequest)
		if responce.StatusCode != http.StatusOK {
			http.Error(w, responce.ErrorMessage, responce.StatusCode)
		}
		json.NewEncoder(w).Encode(responce)
	}
	if r.Method == http.MethodGet {
		responce := users.GetUser(facebookUserId)
		if responce.StatusCode != http.StatusOK {
			http.Error(w, "User Not Found", responce.StatusCode)
		}
		json.NewEncoder(w).Encode(responce)
	}
}
