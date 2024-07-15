package sql

import "time"

type PatientInfo struct {
	UserID      *string    `json:"userID,omitempty" required:"false"`
	FirstName   *string    `json:"firstName" doc:"First name of the user"`
	LastName    *string    `json:"lastName" doc:"Last name of the user" required:"false"`
	Age         *uint16    `json:"age" doc:"Age of the user"`
	BloodGroup  *string    `json:"bloodGroup" doc:"Blood Group of the user" required:"false"`
	PhoneNumber *string    `json:"phoneNumber" doc:"Phone number of the user" minLength:"10" maxLength:"12"`
	Address     *string    `json:"address" doc:"Address of the user" required:"false"`
	DateOfBirth *time.Time `json:"dateOfBirth" doc:"Date of birth of the user" required:"false" format:"date-time"`
}
