package models

type Assignee struct {
	ID 				string `json:"_id" bson:"_id"`
	Name			string `json:"name" bson:"name"`
	Department  	string `json:"department" bson:"department"`
}
