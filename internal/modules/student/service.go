package student

import (
	"context"
	"errors"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateStudent(ctx context.Context, student *Student) error {
	if student.Name == "" || student.Phone == "" {
		return errors.New("name and phone required")
	}

	student.Status = "active"
	student.CheckIn = time.Now()

	return s.repo.Create(ctx, student)
}

func (s *Service) GetAllStudents(ctx context.Context, tenantID string) ([]Student, error) {
	return s.repo.FindAll(ctx, tenantID)
}

func (s *Service) GetStudent(ctx context.Context, tenantID, id string) (*Student, error) {
	return s.repo.FindByID(ctx, tenantID, id)
}

func (s *Service) UpdateStudent(ctx context.Context, tenantID, id string, update map[string]interface{}) error {
	return s.repo.Update(ctx, tenantID, id, update)
}

func (s *Service) DeleteStudent(ctx context.Context, tenantID, id string) error {
	return s.repo.Delete(ctx, tenantID, id)
}
