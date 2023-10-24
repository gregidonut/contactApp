package contact

type Contact struct {
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	PhoneNumber  string `json:"phoneNumber,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
}

func NewContact(firstName, lastName, phoneNumber, email string) (*Contact, error) {
	payload := new(Contact)
	payload.FirstName = firstName
	payload.LastName = lastName
	payload.PhoneNumber = phoneNumber
	payload.EmailAddress = email

	return payload, nil
}
