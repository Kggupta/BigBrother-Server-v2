package models

type Stock struct {
	id		   string `bson:"id,omitempty" json:"id,omitempty"`
	priceArray []Price `bson:"prices,omitempty" json:"prices,omitempty"`
}