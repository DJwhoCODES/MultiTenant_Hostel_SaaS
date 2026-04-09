package student

import (
	"context"
	"time"

	"github.com/djwhocodes/hostel_saas/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(ctx context.Context, student *Student) error {
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	_, err := database.DB.Collection("students").InsertOne(ctx, student)
	return err
}

func (r *Repository) FindAll(ctx context.Context, tenantID string) ([]Student, error) {
	cursor, err := database.DB.Collection("students").
		Find(ctx, bson.M{"tenantId": tenantID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []Student
	for cursor.Next(ctx) {
		var s Student
		cursor.Decode(&s)
		students = append(students, s)
	}

	return students, nil
}

func (r *Repository) FindByID(ctx context.Context, tenantID, id string) (*Student, error) {
	objID, _ := primitive.ObjectIDFromHex(id)

	var student Student
	err := database.DB.Collection("students").
		FindOne(ctx, bson.M{
			"_id":      objID,
			"tenantId": tenantID,
		}).
		Decode(&student)

	return &student, err
}

func (r *Repository) Update(ctx context.Context, tenantID, id string, update bson.M) error {
	objID, _ := primitive.ObjectIDFromHex(id)

	update["updatedAt"] = time.Now()

	_, err := database.DB.Collection("students").
		UpdateOne(ctx,
			bson.M{"_id": objID, "tenantId": tenantID},
			bson.M{"$set": update},
		)

	return err
}

func (r *Repository) Delete(ctx context.Context, tenantID, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := database.DB.Collection("students").
		DeleteOne(ctx, bson.M{
			"_id":      objID,
			"tenantId": tenantID,
		})

	return err
}
