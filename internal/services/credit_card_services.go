package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"my-budget-planner/internal/postgres/models"
	"my-budget-planner/internal/repository"
)

type CreditCardService struct {
	repo *repository.CreditCardRepository
}

func NewCreditCardService(pool *pgxpool.Pool) *CreditCardService {
	return &CreditCardService{repo: repository.NewCreditCardRepository(pool)}
}

func (s *CreditCardService) GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]models.CreditCard, error) {
	return s.repo.FetchAllByUserID(ctx, userID)
}

func (s *CreditCardService) GetByID(ctx context.Context, id uuid.UUID) (*models.CreditCard, error) {
	return s.repo.FetchOneByID(ctx, id)
}

func (s *CreditCardService) Create(ctx context.Context, cc *models.CreditCard) error {
	return s.repo.Create(ctx, cc)
}

func (s *CreditCardService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
