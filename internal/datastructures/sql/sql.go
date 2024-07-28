package sql

import (
	"fmt"
	"strings"
	"time"
)

type Date string

const ctLayout = "2006-01-02"

// UnmarshalJSON Parses the json string in the custom format
func (ct *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	fmt.Println("tt", s)
	if _, err = time.Parse(ctLayout, s); err != nil {
		return err
	}
	fmt.Println("error: ", err)
	*ct = Date(s)
	fmt.Println("thing ", *ct)
	return nil
}

// MarshalJSON writes a quoted string in the custom format
func (ct Date) MarshalJSON() ([]byte, error) {
	return []byte(ct), nil
}

type PatientInfo struct {
	RegistrationDate Date    `json:"registrationDate" doc:"Registration date of the user" required:"false" format:"date"`
	Type             *string `json:"type" doc:"Type of the user" enum:"Mr.,Mrs.,Dr." required:"false"`
	Gender           *string `json:"gender" doc:"Gender of the user" enum:"Male,Female" required:"false"`
	ReferedBy        *string `json:"referedBy" doc:"Refered by of the user" required:"false"`
	Email            *string `json:"email" doc:"Email address of the user" required:"false"`
	UserID           *string `json:"userID,omitempty" required:"false"`
	FirstName        *string `json:"firstName" doc:"First name of the user"`
	LastName         *string `json:"lastName" doc:"Last name of the user" required:"false"`
	Age              *uint16 `json:"age" doc:"Age of the user" required:"false"`
	BloodGroup       *string `json:"bloodGroup" doc:"Blood Group of the user" required:"false"`
	PhoneNumber      *string `json:"phoneNumber" doc:"Phone number of the user" minLength:"10" maxLength:"12" required:"false"`
	Address          *string `json:"address" doc:"Address of the user" required:"false"`
	PostCode         *int64  `json:"postCode" doc:"PostCode of the user" required:"false"`
	State            *string `json:"state" doc:"State of the user" required:"false"`
	City             *string `json:"city" doc:"City of the user" required:"false"`
	DateOfBirth      Date    `json:"dateOfBirth" doc:"Date of birth of the user" required:"false" format:"date"`
}

type Allergies struct {
	FoodAllergies   *string `json:"foodAllergies" doc:"Food Allergies of the user" required:"false"`
	DrugAllergies   *string `json:"drugAllergies" doc:"Drug Allergies of the user" required:"false"`
	AirboneAlergies *string `json:"airboneAllergies" doc:"Airbone Allergies of the user" required:"false"`
}

type MedicalHistory struct {
	MedHistory *string `json:"medicalHistory" doc:"Medical Conditions of the user" required:"false"`
}

type HospitalizationHistory struct {
	HosHistory *string `json:"hospitalizationHistory" doc:"Current Hospitalization of the user" required:"false"`
}
