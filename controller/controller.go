package controller

import (
	"context"
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

// ReadCandidates godoc
// @Summary Read candidates
// @Description Read all candidates
// @Accept  json
// @Produce  json
// @Param id query string false "name search by id"
// @Success 200 {object} object model.Account
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /readCandidate [get]
func (c *Controller) ReadCandidates(ctx *gin.Context) {
	id := ctx.Query("id")
	var candidate models.Candidate
	c.Database.Collection(models.CandidateTableName()).FindOne(ctx, bson.M{"_id":id}).Decode(&candidate)
	ctx.JSON(http.StatusOK, candidate)
}
