package router

import (
	"encoding/json"
	"fmt"
	"godoc/pkg/datastructures"
	"godoc/pkg/db"
	"log"
	"net/http"
)

func HandleUserData(w http.ResponseWriter, r *http.Request) {
	// Initialize storage
	storage := db.NewJSONFileStorage("./data")

	switch r.Method {
	case http.MethodPost:
		var userData datastructures.PatientData
		if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Validate required fields
		if userData.FirstName == "" || userData.MobileNo == "" {
			http.Error(w, "Username and mobile number are required", http.StatusBadRequest)
			return
		}

		// Save to storage
		if err := storage.Save(userData.FirstName, userData.MobileNo, userData); err != nil {
			log.Printf("Failed to save user data: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		fmt.Printf("Received and saved user data: %+v\n", userData)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
		return
	case http.MethodGet:
		// Get query parameters
		firstName := r.URL.Query().Get("firstname")
		mobileNo := r.URL.Query().Get("mobileno")

		// Validate required parameters
		if firstName == "" || mobileNo == "" {
			http.Error(w, "firstname and mobileno parameters are required", http.StatusBadRequest)
			return
		}

		// Check if data exists
		if !storage.Exists(firstName, mobileNo) {
			http.Error(w, "User data not found", http.StatusNotFound)
			return
		}

		// Load data from storage
		var userData datastructures.PatientData
		if err := storage.Load(firstName, mobileNo, &userData); err != nil {
			log.Printf("Failed to load user data: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Return data as JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userData)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Initialize storage
	storage := db.NewJSONFileStorage("./data")

	// Get list of all patients
	patients := storage.List()

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Return the list as JSON
	if err := json.NewEncoder(w).Encode(patients); err != nil {
		log.Printf("Failed to encode patient list: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
