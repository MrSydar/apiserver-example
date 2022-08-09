package models

type Note struct {
	Protected `bson:",inline"`

	Text string `bson:"text" json:"text"`
}
