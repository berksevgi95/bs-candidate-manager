package models

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	
	"../utils"
)

type Candidate struct {
	ID 				string `json:"_id" bson:"_id"`
	FirstName 		string `json:"first_name" bson:"first_name"`
	LastName  		string `json:"last_name" bson:"last_name"`
	Email  			string `json:"email" bson:"email"`
	Department  	string `json:"department" bson:"department"`
	University  	string `json:"university" bson:"university"`
	Experience  	bool `json:"experience" bson:"experience"`
	Status			string `json:"status" bson:"status"`
	MeetingCount 	int `json:"meeting_count" bson:"meeting_count"`
	NextMeeting 	primitive.DateTime `json:"next_meeting" bson:"next_meeting"`
	Assignee 		string `json:"assignee" bson:"assignee"`
	ApplicationDate	primitive.DateTime `json:"application_date,omitempty" bson:"application_date,omitempty"`
}

func (c *Candidate) Deny() (_candidate *Candidate) {
	c.Status = "Denied";
	return c;
}

func (c *Candidate) Accept() (_candidate *Candidate) {
	c.Status = "Accepted";
	return c;
}


func ReadCandidate(ctx *gin.Context, db *mongo.Database, m map[string]string) (*Candidate, error) {
	var candidate Candidate
	dbErr := db.Collection("Candidates").FindOne(ctx, bson.M{"_id":m["id"]}).Decode(&candidate)
	if dbErr != nil {
		return nil, dbErr
	}
	return &candidate, nil
}

func CreateCandidate(ctx *gin.Context, db *mongo.Database, m map[string]string) (*Candidate, error) {
	experience, experienceError := strconv.ParseBool(m["experience"])
	if experienceError != nil {
		return nil, errors.New("Could not convert experience")
	}

	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(m["email"]) {
		return nil, errors.New("Not valid email")
	}

	if utils.DepartmentExist(m["department"]) == -1 {
		return nil, errors.New("Not valid department")
	}

	foundAssignee, _ := FindAssigneeByID(ctx, db, map[string]string{ "id": m["assignee"]})
	if foundAssignee == nil {
		return nil, errors.New("Assignee not found")	
	}

	if foundAssignee.Department != m["department"] {
		return nil, errors.New("Not related department")	
	}

	createdCandidate := &Candidate{
		ID: primitive.NewObjectID().Hex(),
		FirstName: m["first_name"],
		LastName: m["last_name"],
		Email: m["email"],
		Department: m["department"],
		University: m["university"],
		Experience: experience,
		Status: "Pending",
		MeetingCount: 0,
		NextMeeting: primitive.NewDateTimeFromTime(time.Now()),
		Assignee: m["assignee"],
		ApplicationDate: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, dbErr := db.Collection("Candidates").InsertOne(ctx, createdCandidate)
	if dbErr != nil {
		return nil, dbErr
	}

	return createdCandidate, nil
}

func DeleteCandidate(ctx *gin.Context, db *mongo.Database, m map[string]string) (int64, error) {
	dbResult, dbErr := db.Collection("Candidates").DeleteOne(ctx, bson.M{"_id":m["id"]})
	if dbErr != nil {
		return 0, dbErr
	}
	return dbResult.DeletedCount, nil
}

func DenyCandidate(ctx *gin.Context, db *mongo.Database, m map[string]string) (*Candidate, error) {
	var candidate Candidate
	db.Collection("Candidates").FindOne(ctx, bson.M{"_id":m["id"]}).Decode(&candidate)
	updateDocument := bson.M {
		"$set": candidate.Deny(),
	}
	_, dbErr := db.Collection("Candidates").UpdateOne(ctx, bson.M{"_id":m["id"]}, updateDocument)
	if dbErr != nil {
		return nil, dbErr
	}
	return &candidate, nil
}

func AcceptCandidate(ctx *gin.Context, db *mongo.Database, m map[string]string) (*Candidate, error) {
	var candidate Candidate
	db.Collection("Candidates").FindOne(ctx, bson.M{"_id":m["id"]}).Decode(&candidate)
	updateDocument := bson.M {
		"$set": candidate.Accept(),
	}
	_, dbErr := db.Collection("Candidates").UpdateOne(ctx, bson.M{"_id":m["id"]}, updateDocument)
	if dbErr != nil {
		return nil, dbErr
	}
	return &candidate, nil
}

func FindAssigneesCandidates(ctx *gin.Context, db *mongo.Database, m map[string]string) (*[]Candidate, error) {
	var candidates []Candidate
	dbResult, dbErr := db.Collection("Candidates").Find(ctx, bson.M{"assignee": m["assignee"]})
	if dbErr != nil {
		return nil, dbErr
	}
	dbResult.All(ctx, &candidates)
	return &candidates, nil
}