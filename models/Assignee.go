// Package models represents model(s) which is used for correlation between DO and BO
// Assignee represents assignees stored in the DB
package models

type Assignee struct {
	ID 				string `json:"_id" bson:"_id"`
	Name			string `json:"name" bson:"name"`
	Department  	string `json:"department" bson:"department"`
}
