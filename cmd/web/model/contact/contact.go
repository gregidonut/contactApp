package contact

type Contact struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	PhoneNumber  string `json:"phoneNumber,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
}

func NewContact(id int, firstName, lastName, phoneNumber, email string) (*Contact, error) {
	payload := new(Contact)
	payload.ID = id
	payload.FirstName = firstName
	payload.LastName = lastName
	payload.PhoneNumber = phoneNumber
	payload.EmailAddress = email

	return payload, nil
}
