package database

import (
	"context"
	"log"
	"time"

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
			Keys: map[string]int{
				"tenantId": 1,
				"roomId":   1,
			},
		},
		{
			Keys: map[string]int{
				"tenantId": 1,
				"phone":    1,
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
			Keys: map[string]int{
				"tenantId": 1,
				"dueDate":  1,
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
			Keys: map[string]int{
				"tenantId": 1,
				"email":    1,
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Println("User index error:", err)
	}
}
