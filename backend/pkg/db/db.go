package db

import (
	"encoding/json"
	"fmt"
	"godoc/pkg/datastructures"
	"os"
	"path/filepath"
)

// Storage interface defines the contract for data storage operations
type Storage interface {
	Save(username, mobileNumber string, data interface{}) error
	Load(username, mobileNumber string, data interface{}) error
	Delete(username, mobileNumber string) error
	Exists(username, mobileNumber string) bool
	List() datastructures.PatientListResponse
}

// JSONFileStorage implements Storage interface using JSON files
type JSONFileStorage struct {
	basePath string
}

// NewJSONFileStorage creates a new JSONFileStorage instance
func NewJSONFileStorage(basePath string) *JSONFileStorage {
	// Ensure the directory exists
	os.MkdirAll(basePath, 0755)
	return &JSONFileStorage{
		basePath: basePath,
	}
}

// getFilePath generates the file path for a given username and mobile number
func (j *JSONFileStorage) getFilePath(username, mobileNumber string) string {
	filename := fmt.Sprintf("%s-%s.json", username, mobileNumber)
	return filepath.Join(j.basePath, filename)
}

// Save stores data in JSON format with filename pattern username-mobilenumber.json
func (j *JSONFileStorage) Save(username, mobileNumber string, data interface{}) error {
	filePath := j.getFilePath(username, mobileNumber)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

// Load reads data from JSON file
func (j *JSONFileStorage) Load(username, mobileNumber string, data interface{}) error {
	filePath := j.getFilePath(username, mobileNumber)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(data); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil
}

// Delete removes the JSON file
func (j *JSONFileStorage) Delete(username, mobileNumber string) error {
	filePath := j.getFilePath(username, mobileNumber)
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}

// Exists checks if the JSON file exists
func (j *JSONFileStorage) Exists(username, mobileNumber string) bool {
	filePath := j.getFilePath(username, mobileNumber)
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// List returns a list of all stored patient records with basic information
func (j *JSONFileStorage) List() []datastructures.PatientListResponse {
	var patients []datastructures.PatientListResponse

	// Read all files in the directory
	files, err := os.ReadDir(j.basePath)
	if err != nil {
		// Return empty slice if directory doesn't exist or can't be read
		return patients
	}

	// Process each JSON file
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		// Load patient data from file
		var patientData datastructures.PatientData
		filePath := filepath.Join(j.basePath, file.Name())

		fileHandle, err := os.Open(filePath)
		if err != nil {
			continue // Skip files that can't be opened
		}

		if err := json.NewDecoder(fileHandle).Decode(&patientData); err != nil {
			fileHandle.Close()
			continue // Skip files that can't be decoded
		}
		fileHandle.Close()

		// Add to response list
		patients = append(patients, datastructures.PatientListResponse{
			FirstName: patientData.FirstName,
			LastName:  patientData.LastName,
			MobileNo:  patientData.MobileNo,
		})
	}

	return patients
}
