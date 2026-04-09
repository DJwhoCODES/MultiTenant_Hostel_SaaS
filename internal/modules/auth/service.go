package auth

import (
	"context"
	"errors"

	"github.com/djwhocodes/hostel_saas/internal/utils"
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Signup(ctx context.Context, name, email, password, hostelName string) (string, error) {

	existing, _ := s.repo.GetUserByEmail(ctx, email)
	if existing != nil && existing.Email != "" {
		return "", errors.New("user already exists")
	}

	tenant := &Tenant{
		ID:        uuid.New().String(),
		Name:      hostelName,
		Subdomain: "",
	}

	if err := s.repo.CreateTenant(ctx, tenant); err != nil {
		return "", err
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user := &User{
		ID:       uuid.New().String(),
		TenantID: tenant.ID,
		Name:     name,
		Email:    email,
		Password: hashed,
		Role:     "owner",
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user.ID, user.TenantID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID, user.TenantID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
