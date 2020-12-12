package models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Assignee struct {
	ID 				string `json:"_id" bson:"_id"`
	Name			string `json:"name" bson:"name"`
	Department  	string `json:"department" bson:"department"`
}

func FindAssigneeIDByName(ctx *gin.Context, db *mongo.Database, m map[string]string) (c *Assignee, err error) {
	id := ctx.Query("id")
	var assignee Assignee
	dbErr := db.Collection("Assignees").FindOne(ctx, bson.M{"_id":id}).Decode(&assignee)
	if dbErr != nil {
		return nil, dbErr
	}
	return &assignee, nil
}
