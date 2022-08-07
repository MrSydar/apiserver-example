package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Email string `bson:"email,omitempty" validate:"required"`
}
