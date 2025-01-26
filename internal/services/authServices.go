package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"my-budget-planner/cmd/app/auth"
	"my-budget-planner/internal/postgres/models"
	"my-budget-planner/internal/repository"
	"time"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(pool *pgxpool.Pool) *AuthService {
	return &AuthService{repo: repository.NewAuthRepository(pool)}
}

// SaveRefreshToken saves the refresh token in the database
func (s *AuthService) SaveRefreshToken(userId, refreshToken string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.repo.StoreRefreshToken(ctx, userId, refreshToken)
}

// ValidateRefreshToken returns the refresh token is valid and has not expired, or an empty token otherwise
func (s *AuthService) ValidateRefreshToken(ctx context.Context, userId, token string) (models.RefreshToken, error) {

	refreshToken, err := s.repo.GetRefreshToken(ctx, token)
	if err != nil {
		return refreshToken, fmt.Errorf("invalid refresh token")
	}

	//checks if the token belongs to the request user
	if refreshToken.UserID != userId {
		return refreshToken, fmt.Errorf("refresh token doesnt belong to requesting user")
	}

	//checks if the token has expired
	if time.Now().After(refreshToken.ExpiresAt) {
		return refreshToken, fmt.Errorf("refresh token has expired")
	}

	return refreshToken, nil
}

// DeleteRefreshToken deletes the refresh token from the database
func (s *AuthService) DeleteRefreshToken(token string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.repo.DeleteRefreshToken(ctx, token)
}

// RefreshToken validates refresh token and refreshes the access token
func (s *AuthService) RefreshToken(userId, token string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	refreshToken, err := s.ValidateRefreshToken(ctx, userId, token)
	if err != nil {
		return "", err
	}

	newAccessToken, err := auth.GenerateAccessToken(refreshToken.UserID)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
