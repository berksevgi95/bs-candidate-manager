package service

import (
	"errors"
	"regexp"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"../models"
	"../repository"
	"../utils"
)


func CreateCandidate(candidate *models.Candidate) (*models.Candidate, error) {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(candidate.Email) {
		return nil, errors.New("Not valid email")
	}

	if utils.DepartmentExist(candidate.Department) == -1 {
		return nil, errors.New("Not valid department")
	}

	foundAssignee, _ := FindAssigneeByID(candidate.Assignee)
	if foundAssignee == nil {
		return nil, errors.New("Assignee not found")	
	}

	if foundAssignee.Department != candidate.Department {
		return nil, errors.New("Not related department")	
	}

	newCandidate := models.Candidate{
		ID: primitive.NewObjectID().Hex(),
		FirstName: candidate.FirstName,
		LastName: candidate.LastName,
		Email: candidate.Email,
		Department: candidate.Department,
		University: candidate.University,
		Experience: candidate.Experience,
		Status: "Pending",
		MeetingCount: 0,
		NextMeeting: primitive.NewDateTimeFromTime(time.Now()),
		Assignee: candidate.Assignee,
		ApplicationDate: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, createError := repository.CreateCandidate(newCandidate)
	if createError != nil {
		return nil, createError
	}
	return &newCandidate, nil
}


func ReadCandidate(_id string) (*models.Candidate, error) {
	readResult, readError := repository.ReadCandidate(bson.M{"_id":_id})
	if readError != nil {
		return nil, readError
	}
	return readResult, nil
}


func DeleteCandidate(_id string) (error) {
	_, deleteError := repository.DeleteCandidate(bson.M{"_id":_id})
	if deleteError != nil {
		return deleteError
	}
	return nil
}


func ArrangeMeeting(_id string, nextMeetingTime *time.Time) (error) {
	candidate, dbError := ReadCandidate(_id)
	if dbError != nil {
		return dbError
	}
	if candidate.MeetingCount == 4 {
		return errors.New("No more meeting")
	} else if candidate.MeetingCount == 3 {
		foundAssignee, _ := FindAssigneeIDByName("Zafer")
		candidate.Assignee = foundAssignee
	}
	candidate.NextMeeting = primitive.NewDateTimeFromTime(*nextMeetingTime)

	_, dbError = repository.UpdateCandidate(candidate)
	if dbError != nil {
		return dbError
	}
	return nil
}


func CompleteMeeting(_id string) (error) {
	candidate, dbError := ReadCandidate(_id)
	
	if candidate.MeetingCount == 4 {
		return errors.New("Not more meeting")
	} else if candidate.MeetingCount == 0 {
		candidate.Status = "In Progress"
	}
	candidate.MeetingCount = candidate.MeetingCount + 1

	_, dbError = repository.UpdateCandidate(candidate)
	if dbError != nil {
		return dbError
	}
	return nil
}


func DenyCandidate(_id string) (error) {
	readResult, readError := ReadCandidate(_id)
	if readError != nil {
		return readError
	}

	if readResult.MeetingCount == 4 || readResult.Status == "In Progress" {
		readResult.Deny()
	} else {
		return errors.New("Not enough meeting")
	}

	_, updateError := repository.UpdateCandidate(readResult)
	if updateError != nil {
		return updateError
	}
	return nil
}



func AcceptCandidate(_id string) (error) {
	readResult, readError := ReadCandidate(_id)
	if readError != nil {
		return readError
	}

	if readResult.MeetingCount == 4 || readResult.Status == "In Progress" {
		readResult.Accept()
	} else {
		return errors.New("Not enough meeting")
	}

	_, readError = repository.UpdateCandidate(readResult)
	if readError != nil {
		return readError
	}
	return nil
}


func FindAssigneeIDByName(name string) (string, error) {
	readResult, readError := repository.ReadAssignee(bson.M{"name":name})
	if readError != nil {
		return "", readError
	}
	return readResult.ID, nil
}


func FindAssigneesCandidates(_id string) (*[]models.Candidate, error) {
	readResult, readError := repository.ReadCandidates(bson.M{"assignee":_id})
	if readError != nil {
		return nil, readError
	}
	return readResult, nil
}


func FindAssigneeByID(_id string) (*models.Assignee, error) {
	readResult, readError := repository.ReadAssignee(bson.M{"_id":_id})
	if readError != nil {
		return nil, readError
	}
	return readResult, nil
}

