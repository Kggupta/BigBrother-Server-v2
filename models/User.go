package models

type User struct {
	id			  string `bson:"id,omitempty" json:"id,omitempty"`
	shackName	  string `bson:"shackName,omitempty" json:"shackName,omitempty"`
	donatorStatus string `bson:"donator,omitempty" json:"donator,omitempty"`
}