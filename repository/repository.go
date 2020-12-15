// Package repository is responsible for accessing DB on CRUD operations
package repository

import (
	"context"
	"time"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../models"
	"../utils"
)

// Register returns mongo.Database
func Register() {
	utils.Singleton(func() *mongo.Database {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		clientOptions := options.Client().ApplyURI("mongodb://" + os.Getenv("REPOHOST") + ":" + os.Getenv("REPOPORT"))
		collection, _ := mongo.Connect(ctx, clientOptions)
		return collection.Database(os.Getenv("REPODB"))
	})
}

// ReadCandidate returns models.Candidate, error
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

// ReadCandidates returns models.Candidate[], error
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

// CreateCandidate returns models.Candidate, error
func CreateCandidate(candidate models.Candidate) (bool, error) {
	var db *mongo.Database
	utils.Make(&db)

	_, error := db.Collection("Candidates").InsertOne(context.Background(), candidate)
	if error != nil {
		return false, error
	}
	return true, nil
}

// DeleteCandidate returns bool, error
func DeleteCandidate(filter primitive.M) (bool, error) {
	var db *mongo.Database
	utils.Make(&db)

	_, error := db.Collection("Candidates").DeleteOne(context.Background(), filter)
	if error != nil {
		return false, error
	}
	return true, nil
}

// UpdateCandidate returns bool, error
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

// ReadAssignee returns models.Assignee, error
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


