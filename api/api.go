package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"../models"
	"../service"
)

// ReadCandidate godoc
// @Summary Read candidate
// @Description Reads a specific candidate by id
// @Accept  json
// @Produce  json
// @Param id query string false "Candidate ID"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /readCandidate [get]
func ReadCandidate(ctx *gin.Context) {
	result, error := service.ReadCandidate(ctx.Query("id"))
	if error != nil {
		ctx.JSON(http.StatusNotFound, error.Error())
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

// CreateCandidate godoc
// @Summary Create candidate
// @Description Creates a new candidate
// @Accept  json
// @Produce  json
// @Param first_name query string true "First Name"
// @Param last_name query string true "Last Name"
// @Param email query string false "E-Mail"
// @Param department query string false "Department"
// @Param university query string false "University"
// @Param experience query boolean false "Experience"
// @Param assignee query string true "Assignee"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /createCandidate [post]
func CreateCandidate(ctx *gin.Context) {
	experience, experienceError := strconv.ParseBool(ctx.Query("experience"))
	if experienceError != nil {
		ctx.JSON(http.StatusNotFound, errors.New("Could not convert experience"))
	}
	result, error := service.CreateCandidate(&models.Candidate{
		FirstName: ctx.Query("first_name"),
		LastName: ctx.Query("last_name"),
		Email: ctx.Query("email"),
		Department: ctx.Query("department"),
		University: ctx.Query("university"),
		Experience: experience,
		Assignee: ctx.Query("assignee"),
	})
	if error != nil {
		ctx.JSON(http.StatusNotFound, error.Error())
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

// DeleteCandidate godoc
// @Summary Delete candidate
// @Description Removes a candidate by id
// @Accept  json
// @Produce  json
// @Param id query string false "Candidate ID"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /deleteCandidate [delete]
func DeleteCandidate(ctx *gin.Context) {
	error := service.DeleteCandidate(ctx.Query("id"))
	if error != nil {
		ctx.JSON(http.StatusNotFound, error.Error())
	} else {
		ctx.JSON(http.StatusOK, true)
	}
}

// DenyCandidate godoc
// @Summary Deny candidate
// @Description Denies a candidate by id
// @Accept  json
// @Produce  json
// @Param id query string false "Candidate ID"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /denyCandidate [put]
func DenyCandidate(ctx *gin.Context) {
	error := service.DenyCandidate(ctx.Query("id"))
	if error != nil {
		ctx.JSON(http.StatusNotFound, error.Error())
	} else {
		ctx.JSON(http.StatusOK, true)
	}
}

// AcceptCandidate godoc
// @Summary Accept candidate
// @Description Accepts a candidate
// @Accept  json
// @Produce  json
// @Param id query string false "Candidate ID"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /acceptCandidate [put]
func AcceptCandidate(ctx *gin.Context) {
	error := service.AcceptCandidate(ctx.Query("id"))
	if error != nil {
		ctx.JSON(http.StatusNotFound, error.Error())
	} else {
		ctx.JSON(http.StatusOK, true)
	}
}

// FindAssigneeIDByName godoc
// @Summary Find assignee ID by name
// @Description Finds ID of an assignee by name
// @Accept  json
// @Produce  json
// @Param name query string false "Assignee Name"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /findAssigneeIDByName [get]
func FindAssigneeIDByName(ctx *gin.Context) {
	result, error := service.FindAssigneeIDByName(ctx.Query("name"))
	if error != nil {
		ctx.JSON(http.StatusNotFound, error.Error())
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

// FindAssigneesCandidates godoc
// @Summary Find asignees candidates
// @Description Finds candidates of an assignee by ID of assignee
// @Accept  json
// @Produce  json
// @Param id query string false "Assignee ID"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /findAssigneesCandidates [get]
func FindAssigneesCandidates(ctx *gin.Context) {
	result, error := service.FindAssigneesCandidates(ctx.Query("id"))
	if error != nil {
		ctx.JSON(http.StatusNotFound, error.Error())
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

// ArrangeMeeting godoc
// @Summary Arrange meeting
// @Description Aranges a meeting with an available candidate
// @Accept  json
// @Produce  json
// @Param id query string false "Candidate ID"
// @Param nextMeetingTime query string false "Next Meeting Time (Ex. 2020-11-11T20:59:48.133+03:00)"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /arrangeMeeting [put]
func ArrangeMeeting(ctx *gin.Context) {
	time, _ := time.Parse(time.RFC3339, ctx.Query("nextMeetingTime"))
	error := service.ArrangeMeeting(ctx.Query("_id"), &time)
	if error != nil {
		ctx.JSON(http.StatusNotFound, error.Error())
	} else {
		ctx.JSON(http.StatusOK, true)
	}
}

// CompleteMeeting godoc
// @Summary Complete meeting
// @Description Completes a meeting of specified candidate
// @Accept  json
// @Produce  json
// @Param id query string false "Candidate ID"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /completeMeeting [put]
func CompleteMeeting(ctx *gin.Context) {
	error := service.CompleteMeeting(ctx.Query("id"))
	if error != nil {
		ctx.JSON(http.StatusNotFound, error.Error())
	} else {
		ctx.JSON(http.StatusOK, true)
	}
}