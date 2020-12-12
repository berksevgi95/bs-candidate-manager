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
	var assignee Assignee
	dbErr := db.Collection("Assignees").FindOne(ctx, bson.M{"name":m["name"]}).Decode(&assignee)
	if dbErr != nil {
		return nil, dbErr
	}
	return &assignee, nil
}
