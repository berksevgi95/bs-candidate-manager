package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Candidate struct {
	ID 				string `json:"_id" bson:"_id"`
	FirstName 		string `json:"first_name" bson:"first_name"`
	LastName  		string `json:"last_name" bson:"last_name"`
	Email  			string `json:"email" bson:"email"`
	Department  	string `json:"department" bson:"department"`
	University  	string `json:"university" bson:"university"`
	Experience  	bool `json:"experience" bson:"experience"`
	Status  		string `json:"status" bson:"status"`
	MeetingCount 	int `json:"meeting_count" bson:"meeting_count"`
	NextMeeting 	primitive.DateTime `json:"next_meeting" bson:"next_meeting"`
	Assignee 		string `json:"assignee" bson:"assignee"`
	ApplicationDate	primitive.DateTime `json:"application_date,omitempty" bson:"application_date,omitempty"`
}

func CandidateTableName() string {
	return "Candidates"
}
