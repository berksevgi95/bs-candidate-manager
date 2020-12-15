package service

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"../models"
	"../repository"


)

func TestReadCandidate(t *testing.T) {
	repository.Register()
	_, readError := ReadCandidate("5b758c6151d9590001def630")
	if readError != nil {
		t.Errorf(readError.Error())
	}
}

func TestCreateCandidate(t *testing.T) {
	repository.Register()
	_, readError := CreateCandidate(&models.Candidate{
		ID: primitive.NewObjectID().Hex(),
		FirstName: "Berk",
		LastName: "Sevgi",
		Email: "berksevgi95@gmail.com",
		Department: "Design",
		University: "Hacettepe",
		Experience: true,
		Status: "Pending",
		MeetingCount: 0,
		NextMeeting: primitive.NewDateTimeFromTime(time.Now()),
		Assignee: "5c18ae31a7948900011168b9",
		ApplicationDate: primitive.NewDateTimeFromTime(time.Now()),
	})
	if readError != nil {
		t.Errorf(readError.Error())
	}
}


func TestDeleteCandidate(t *testing.T) {
	repository.Register()
	readError := DeleteCandidate("5fd91cc99467c7b34a8e300b")
	if readError != nil {
		t.Errorf(readError.Error())
	}
}