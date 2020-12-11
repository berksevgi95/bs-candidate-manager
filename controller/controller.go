package controller

import (
	"net/http"
	"context"
	"time"

	"../models"
	"../config"
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
// @Param _id path string true "Account ID"
// @Accept json
// @Produce json
// @Success 200 {array} object models.Candidate
// @Header 200 {string} Token "qwerty"
// @Failure default {object} object httputil.DefaultError
// @Router /candidates/read [get]
func (c *Controller) ReadCandidates(ctx *gin.Context) {
	var candidates []models.Candidate
	cursor, err := c.Database.Collection(models.CandidateTableName()).Find(ctx, bson.M{})
	if err != nil {
		// response.WriteHeader(http.StatusInternalServerError)
		// response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var candidate models.Candidate
		cursor.Decode(&candidate)
		candidates = append(candidates, candidate)
	}
	if err := cursor.Err(); err != nil {
		// response.WriteHeader(http.StatusInternalServerError)
		// response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx.JSON(http.StatusOK, candidates)
}
