package usecases

import (
	"context"
	"talana_prueba_tecnica/src/entity/requests"
	"talana_prueba_tecnica/src/entity/responses"
)

type UserUseCaseInterface interface {
	FindAll(ctx context.Context) ([]responses.UserResponse, error)
	GetUserByID(ctx context.Context, id uint) (responses.UserResponse, error)
	CreateUser(ctx context.Context, user requests.RegisterUserRequest) error
	UpdateUser(ctx context.Context, id uint, user requests.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id uint) error
}
