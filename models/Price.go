package models

type Price struct {
	value int `bson:"value,omitempty" json:"value,omitempty"`
	date int64 `bson:"date,omitempty" json:"date,omitempty"`
}