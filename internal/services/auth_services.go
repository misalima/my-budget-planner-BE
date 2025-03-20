package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"my-budget-planner/cmd/app/auth"
	"my-budget-planner/internal/postgres/models"
	"my-budget-planner/internal/repository"
	"time"
)

type AuthService struct {
	authRepo *repository.AuthRepository
	userRepo *repository.UserRepository
}

func NewAuthService(pool *pgxpool.Pool) *AuthService {
	return &AuthService{
		authRepo: repository.NewAuthRepository(pool),
		userRepo: repository.NewUserRepository(pool),
	}
}

// Login logs the user generating a new access token
func (s *AuthService) Login(email, password string) (string, string, error) {
	//check if the user exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", fmt.Errorf("invalid email or password")
	}
	fmt.Printf("user email: %s \n user first_name: %s \n, user password: %s\n", user.Email, user.FirstName, user.Password)

	//check if the password is correct
	err = s.CheckPasswords(password, user.Password)
	if err != nil {
		return "", "", err
	}

	//generate access token
	accessToken, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}

	//generate refresh token
	refreshToken, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	//save refresh token
	err = s.SaveRefreshToken(user.ID, refreshToken)
	if err != nil {
		return "", "", err
	}

	//return access token and refresh token
	return accessToken, refreshToken, nil

}

// SaveRefreshToken saves the refresh token in the database
func (s *AuthService) SaveRefreshToken(userId uuid.UUID, refreshToken string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.authRepo.StoreRefreshToken(ctx, userId, refreshToken)
}

// ValidateRefreshToken returns the refresh token is valid and has not expired, or an empty token otherwise
func (s *AuthService) ValidateRefreshToken(ctx context.Context, userId uuid.UUID, token string) (models.RefreshToken, error) {

	refreshToken, err := s.authRepo.GetRefreshToken(ctx, token)
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

	return s.authRepo.DeleteRefreshToken(ctx, token)
}

// RefreshToken validates refresh token and refreshes the access token
func (s *AuthService) RefreshToken(userId uuid.UUID, token string) (string, error) {
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

func (s *AuthService) CheckPasswords(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid email or password")
	}
	return nil
}
