package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Protected struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id"`
	Owner primitive.ObjectID `bson:"owner,omitempty" json:"owner"`
}
