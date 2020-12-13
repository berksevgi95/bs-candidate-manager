package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../config"
	"../models"
	"../utils"
)

func Register() {
	utils.Singleton(func() *mongo.Database {
		var cfg = config.GetConfig()
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		clientOptions := options.Client().ApplyURI("mongodb://" + cfg.Host + ":" + cfg.Port)
		collection, _ := mongo.Connect(ctx, clientOptions)
		return collection.Database(cfg.Database)
	})
}


func ReadCandidate(filter primitive.M) (*models.Candidate, error) {
	var db *mongo.Database
	utils.Make(&db)

	var candidate models.Candidate
	error := db.Collection("Candidates").FindOne(context.Background(), filter).Decode(&candidate)
	if error != nil {
		return nil, error
	}
	return &candidate, nil
}

func ReadCandidates(filter primitive.M) (*[]models.Candidate, error) {
	var db *mongo.Database
	utils.Make(&db)

	result, error := db.Collection("Candidates").Find(context.Background(), filter)
	if error != nil {
		return nil, error
	}
	var candidates []models.Candidate
	result.All(context.Background(), &candidates)
	return &candidates, nil
}

func CreateCandidate(candidate models.Candidate) (bool, error) {
	var db *mongo.Database
	utils.Make(&db)

	_, error := db.Collection("Candidates").InsertOne(context.Background(), candidate)
	if error != nil {
		return false, error
	}
	return true, nil
}


func DeleteCandidate(filter primitive.M) (bool, error) {
	var db *mongo.Database
	utils.Make(&db)

	_, error := db.Collection("Candidates").DeleteOne(context.Background(), filter)
	if error != nil {
		return false, error
	}
	return true, nil
}


func UpdateCandidate(candidate *models.Candidate) (bool, error) {
	var db *mongo.Database
	utils.Make(&db)

	update := bson.M {
		"$set": candidate,
	}
	_, error := db.Collection("Candidates").UpdateOne(context.Background(), bson.M{"_id":candidate.ID}, update)
	if error != nil {
		return false, error
	}
	return true, nil
}


func ReadAssignee(filter primitive.M) (*models.Assignee, error) {
	var db *mongo.Database
	utils.Make(&db)

	var candidate models.Assignee
	error := db.Collection("Assignees").FindOne(context.Background(), filter).Decode(&candidate)
	if error != nil {
		return nil, error
	}
	return &candidate, nil
}


