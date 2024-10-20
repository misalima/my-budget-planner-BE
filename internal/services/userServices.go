package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"my-budget-planner/internal/postgres/models"
	"my-budget-planner/internal/repository"
	"time"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	return &UserService{repo: repository.NewUserRepository(pool)}
}

func (s *UserService) RegisterUser(user *models.User) error {
	//Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//Check if the user with the given email already exists
	existingUser, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("user with the email already exists")
	}
	//Proceed to create the user
	return s.repo.CreateUser(ctx, user)
}
