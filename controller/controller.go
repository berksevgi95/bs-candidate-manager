package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"../config"
	"../models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
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
	id := ctx.Query("id")
	var candidate models.Candidate
	c.Database.Collection(models.CandidateTableName()).FindOne(ctx, bson.M{"_id":id}).Decode(&candidate)
	ctx.JSON(http.StatusOK, candidate)
}

// CreateCandidate godoc
// @Summary Create candidate
// @Description Create new candidate
// @Accept  json
// @Produce  json
// @Param id body string false "name search by id"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /createCandidate [post]
func (c *Controller) CreateCandidate(ctx *gin.Context) {
	var candidate models.Candidate
	json.NewDecoder(ctx.Request.Body).Decode(&candidate)
	result, err := c.Database.Collection(models.CandidateTableName()).InsertOne(ctx, candidate)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, result.InsertedID)
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
	id := ctx.Query("id")
	result, err := c.Database.Collection(models.CandidateTableName()).DeleteOne(ctx, bson.M{"_id":id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, result.DeletedCount)
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
	id := ctx.Query("id")
	var candidate models.Candidate
	c.Database.Collection(models.CandidateTableName()).FindOne(ctx, bson.M{"_id":id}).Decode(&candidate)
	update := bson.M {
		"$set": candidate.Deny(),
	}
	result, err := c.Database.Collection(models.CandidateTableName()).UpdateOne(ctx, bson.M{"_id":id}, update)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, result)
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
	id := ctx.Query("id")
	var candidate models.Candidate
	c.Database.Collection(models.CandidateTableName()).FindOne(ctx, bson.M{"_id":id}).Decode(&candidate)
	update := bson.M {
		"$set": candidate.Accept(),
	}
	result, err := c.Database.Collection(models.CandidateTableName()).UpdateOne(ctx, bson.M{"_id":id}, update)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, result)
}

// FindAssigneeIDByName godoc
// @Summary Read candidates
// @Description Read all candidates
// @Accept  json
// @Produce  json
// @Param id query string false "name search by id"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /findAssigneeIDByName [get]
func (c *Controller) FindAssigneeIDByName(ctx *gin.Context) {
	id := ctx.Query("id")
	var assignee models.Assignee
	c.Database.Collection(models.AssigneeTableName()).FindOne(ctx, bson.M{"_id":id}).Decode(&assignee)
	ctx.JSON(http.StatusOK, assignee.Name)
}