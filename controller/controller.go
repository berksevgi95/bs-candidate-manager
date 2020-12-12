package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../config"
	"../models"
)

type Controller struct { 
	Database *mongo.Database
}

func NewController() *Controller {
	var c = &Controller{}
	var cfg = config.GetConfig()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://" + cfg.Host + ":" + cfg.Port)
	collection, _ := mongo.Connect(ctx, clientOptions)
	c.Database = collection.Database(cfg.Database)
	return c;
}

// ReadCandidate godoc
// @Summary Read candidates
// @Description Read all candidates
// @Accept  json
// @Produce  json
// @Param id query string false "name search by id"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /readCandidate [get]
func (c *Controller) ReadCandidate(ctx *gin.Context) {
	readCandidate, creationError := models.ReadCandidate(ctx, c.Database, map[string]string{
		"id": ctx.Query("id"),
	})
	if creationError != nil {
		ctx.JSON(http.StatusNotFound, creationError.Error())
	} else {
		ctx.JSON(http.StatusOK, readCandidate)
	}
}

// CreateCandidate godoc
// @Summary Create candidate
// @Description Create new candidate
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
func (c *Controller) CreateCandidate(ctx *gin.Context) {
	createdCandidate, creationError := models.CreateCandidate(ctx, c.Database, map[string]string{
		"first_name": ctx.Query("first_name"),
		"last_name": ctx.Query("last_name"),
		"email": ctx.Query("email"),
		"department": ctx.Query("department"),
		"university": ctx.Query("university"),
		"experience": ctx.Query("experience"),
		"assignee": ctx.Query("assignee"),
	})
	if creationError != nil {
		ctx.JSON(http.StatusNotFound, creationError.Error())
	} else {
		ctx.JSON(http.StatusOK, createdCandidate)
	}
}

// DeleteCandidate godoc
// @Summary Delete candidates
// @Description Removes a candidate
// @Accept  json
// @Produce  json
// @Param id query string false "name search by id"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /deleteCandidate [delete]
func (c *Controller) DeleteCandidate(ctx *gin.Context) {
	deletedCandidate, creationError := models.DeleteCandidate(ctx, c.Database, map[string]string{
		"id": ctx.Query("id"),
	})
	if creationError != nil {
		ctx.JSON(http.StatusNotFound, creationError.Error())
	} else {
		ctx.JSON(http.StatusOK, deletedCandidate)
	}
}

// DenyCandidate godoc
// @Summary Deny candidates
// @Description Denies a candidate
// @Accept  json
// @Produce  json
// @Param id query string false "name search by id"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /denyCandidate [put]
func (c *Controller) DenyCandidate(ctx *gin.Context) {
	deniedCandidate, creationError := models.DenyCandidate(ctx, c.Database, map[string]string{
		"id": ctx.Query("id"),
	})
	if creationError != nil {
		ctx.JSON(http.StatusNotFound, creationError.Error())
	} else {
		ctx.JSON(http.StatusOK, deniedCandidate)
	}
}

// AcceptCandidate godoc
// @Summary Accept candidates
// @Description Accepts a candidate
// @Accept  json
// @Produce  json
// @Param id query string false "name search by id"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /acceptCandidate [put]
func (c *Controller) AcceptCandidate(ctx *gin.Context) {
	acceptedCandidate, creationError := models.AcceptCandidate(ctx, c.Database, map[string]string{
		"id": ctx.Query("id"),
	})
	if creationError != nil {
		ctx.JSON(http.StatusNotFound, creationError.Error())
	} else {
		ctx.JSON(http.StatusOK, acceptedCandidate)
	}
}

// FindAssigneeIDByName godoc
// @Summary Read candidates
// @Description Read all candidates
// @Accept  json
// @Produce  json
// @Param name query string false "name"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /findAssigneeIDByName [get]
func (c *Controller) FindAssigneeIDByName(ctx *gin.Context) {
	foundAssignee, creationError := models.FindAssigneeIDByName(ctx, c.Database, map[string]string{
		"name": ctx.Query("name"),
	})
	if creationError != nil {
		ctx.JSON(http.StatusNotFound, creationError.Error())
	} else {
		ctx.JSON(http.StatusOK, foundAssignee.ID)
	}
}

// FindAssigneesCandidates godoc
// @Summary Read candidates
// @Description Read all candidates
// @Accept  json
// @Produce  json
// @Param id query string false "name search by id"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /findAssigneesCandidates [get]
func (c *Controller) FindAssigneesCandidates(ctx *gin.Context) {
	foundCandidates, creationError := models.FindAssigneesCandidates(ctx, c.Database, map[string]string{
		"assignee": ctx.Query("id"),
	})
	if creationError != nil {
		ctx.JSON(http.StatusNotFound, creationError.Error())
	} else {
		ctx.JSON(http.StatusOK, foundCandidates)
	}
}