package contact

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName    string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName     string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	PhoneNumber  string             `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	EmailAddress string             `json:"emailAddress,omitempty" bson:"emailAddress,omitempty"`
}
