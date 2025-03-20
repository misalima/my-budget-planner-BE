package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
	"my-budget-planner/internal/postgres/models"
	"my-budget-planner/internal/repository"
	"regexp"
	"time"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	return &UserService{repo: repository.NewUserRepository(pool)}
}

func (s *UserService) RegisterUser(user *models.User) error {

	err := ValidateUser(user)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existingUser, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("user with the email already exists")
	}

	return s.repo.CreateUser(ctx, user)
}

func ValidateUser(user *models.User) error {

	if len(user.Username) < 3 {
		return fmt.Errorf("username must be at least 3 characters")
	}

	if user.Email == "" || !isValidEmail(user.Email) {
		return fmt.Errorf("invalid email address")
	}

	if len(user.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	hasLetter := false
	hasNumber := false
	for _, c := range user.Password {
		switch {
		case 'a' <= c && c <= 'z', 'A' <= c && c <= 'Z':
			hasLetter = true
		case '0' <= c && c <= '9':
			hasNumber = true
		}
	}

	if !hasLetter {
		return fmt.Errorf("password must contain at least one letter")
	}

	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}

	return nil
}

func isValidEmail(email string) bool {
	// A very simple email validation check
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
