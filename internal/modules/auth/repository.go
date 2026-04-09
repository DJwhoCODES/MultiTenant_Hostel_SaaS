package auth

import (
	"context"
	"time"

	"github.com/djwhocodes/hostel_saas/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateTenant(ctx context.Context, tenant *Tenant) error {
	tenant.CreatedAt = time.Now()
	_, err := database.DB.Collection("tenants").InsertOne(ctx, tenant)
	return err
}

func (r *Repository) CreateUser(ctx context.Context, user *User) error {
	user.CreatedAt = time.Now()
	_, err := database.DB.Collection("users").InsertOne(ctx, user)
	return err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := database.DB.Collection("users").
		FindOne(ctx, bson.M{"email": email}).
		Decode(&user)

	return &user, err
}
