package models

type Department string

const (
	Marketing Department = "Marketing"
	Design Department = "Design"
	Development Department = "Development"
)

type Assignee struct {
	ID 				string `json:"_id" bson:"_id"`
	Name			string `json:"name" bson:"name"`
	Department  	Department `json:"department" bson:"department"`
}

func AssigneeTableName() string {
	return "Assignees"
}
