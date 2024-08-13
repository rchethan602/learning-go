package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BookName    string             `json:"bookname,omitempty"`
	Author      string             `json:"author,omitempty"`
	Publication string             `json:"publication,omitempty"`
}
