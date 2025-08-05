package main

import (
	"encoding/json"
	"fmt"
	"godoc/pkg/db"
	"log"
	"net/http"
)

type UserData struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	MobileNumber string `json:"mobile_number"`
}

var storage db.Storage

func handleUserData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var userData UserData
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if userData.Username == "" || userData.MobileNumber == "" {
		http.Error(w, "Username and mobile number are required", http.StatusBadRequest)
		return
	}

	// Save to storage
	if err := storage.Save(userData.Username, userData.MobileNumber, userData); err != nil {
		log.Printf("Failed to save user data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Received and saved user data: %+v\n", userData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func main() {
	// Initialize storage
	storage = db.NewJSONFileStorage("./data")

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/userdata", handleUserData)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
