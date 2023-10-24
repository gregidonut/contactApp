package contact

type Contact struct {
	firstName    string
	lastName     string
	phoneNumber  string
	emailAddress string
}

func NewContact(firstName, lastName, phoneNumber, email string) (*Contact, error) {
	payload := new(Contact)
	payload.firstName = firstName
	payload.lastName = lastName
	payload.phoneNumber = phoneNumber
	payload.emailAddress = email

	return payload, nil
}
