package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createStudentIndexes(ctx)
	createPaymentIndexes(ctx)
	createUserIndexes(ctx)
}

func createStudentIndexes(ctx context.Context) {
	collection := DB.Collection("students")

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "tenantId", Value: 1},
				{Key: "roomId", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "tenantId", Value: 1},
				{Key: "phone", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Println("Student index error:", err)
	}
}

func createPaymentIndexes(ctx context.Context) {
	collection := DB.Collection("payments")

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "tenantId", Value: 1},
				{Key: "dueDate", Value: 1},
			},
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Println("Payment index error:", err)
	}
}

func createUserIndexes(ctx context.Context) {
	collection := DB.Collection("users")

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "tenantId", Value: 1},
				{Key: "email", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Println("User index error:", err)
	}
}
