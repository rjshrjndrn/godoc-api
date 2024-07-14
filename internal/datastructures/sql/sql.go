package sql

type CreatePatient struct {
	FirstName   *string `json:"firstName" doc:"First name of the user"`
	LastName    *string `json:"lastName" doc:"Last name of the user" required:"false"`
	Age         *int    `json:"age" doc:"Age of the user"`
	BloodGroup  *string `json:"bloodGroup" doc:"Blood Group of the user" required:"false"`
	PhoneNumber *uint64 `json:"phoneNumber" doc:"Phone number of the user" required:"false"`
	Address     *string `json:"address" doc:"Address of the user" required:"false"`
	DateOfBirth *string `json:"dateOfBirth" doc:"Date of birth of the user" required:"false" format:"date"`
}
