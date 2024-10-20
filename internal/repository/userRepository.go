package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"my-budget-planner/internal/postgres/models"
	"time"
)

// UserRepository is a struct that defines the repository for the user
type UserRepository struct {
	Conn *pgxpool.Pool
}

func NewUserRepository(Conn *pgxpool.Pool) *UserRepository {
	return &UserRepository{Conn: Conn}
}

func (u UserRepository) CreateUser(ctx context.Context, user *models.User) error {

	sql := `INSERT INTO users (username, first_name, last_name, email, password_hash, income, expenditure_limit, created_at, updated_at) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	log.Print("Executing query")
	tag, err := u.Conn.Exec(ctx, sql, user.Username, user.FirstName, user.LastName, user.Email, user.Password, 0.0, 0.0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("no rows were inserted")
	}

	return nil
}

func (u UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	sql := `SELECT * FROM users WHERE email = $1`
	user := &models.User{}
	err := u.Conn.QueryRow(context.Background(), sql, email).Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.ProfilePicture, &user.Income, &user.ExpenditureLimit, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
