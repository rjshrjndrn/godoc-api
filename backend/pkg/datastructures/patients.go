package datastructures

type PatientData struct {
	Title      string `json:"title"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Gender     string `json:"gender"`
	BloodGroup string `json:"bloodGroup"`
	RegDate    string `json:"regDate"`
	DOB        string `json:"dob"`
	Age        string `json:"age"`
	MobileNo   string `json:"mobileNo"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	PostCode   string `json:"postCode"`
	State      string `json:"state"`
	City       string `json:"city"`
}

type PatientListResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	MobileNo  string `json:"mobileNo"`
}
