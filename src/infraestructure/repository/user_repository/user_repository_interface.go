package repository

import (
	"context"
	"talana_prueba_tecnica/src/entity/models"
)

type UserRepositoryInterface interface {
	FindAll(ctx context.Context) ([]models.UserModel, error)
	FindByID(ctx context.Context, id uint) (*models.UserModel, error)
	Create(ctx context.Context, user *models.UserModel) error
	Update(ctx context.Context, user *models.UserModel, id uint) error
	Delete(ctx context.Context, id uint) error
}
