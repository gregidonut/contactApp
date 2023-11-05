package contact

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName    string             `json:"firstName,omitempty"`
	LastName     string             `json:"lastName,omitempty"`
	PhoneNumber  string             `json:"phoneNumber,omitempty"`
	EmailAddress string             `json:"emailAddress,omitempty"`
}
