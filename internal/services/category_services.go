package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"my-budget-planner/internal/postgres/models"
	"my-budget-planner/internal/repository"
	"time"
)

type CategoryServices struct {
	repo *repository.CategoryRepository
}

func NewCategoryServices(db *pgxpool.Pool) *CategoryServices {
	return &CategoryServices{repo: repository.NewCategoryRepository(db)}
}

func (c *CategoryServices) CreateCategory(category *models.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := c.repo.CheckUserExists(ctx, category.UserID)
	if err != nil {
		return err
	}
	return c.repo.CreateCategory(ctx, category)
}

func (c *CategoryServices) GetCategoriesByUserID(userId uuid.UUID) ([]models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.repo.GetCategoryByUserID(ctx, userId)
}

func (c *CategoryServices) DeleteCategory(categoryId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.repo.DeleteCategory(ctx, categoryId)
}
