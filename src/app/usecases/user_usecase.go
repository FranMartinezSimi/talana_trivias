package usecases

import (
	"context"
	"talana_prueba_tecnica/src/entity/models"
	"talana_prueba_tecnica/src/entity/requests"
	"talana_prueba_tecnica/src/entity/responses"
	"talana_prueba_tecnica/src/infraestructure/repository"

	"github.com/sirupsen/logrus"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (u *UserUseCase) FindAll(ctx context.Context) ([]responses.UserResponse, error) {
	log := logrus.WithContext(ctx)
	log.Info("Get all users usecase")

	result, err := u.repository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("Users found")
	var usersList []responses.UserResponse

	for _, users := range result {
		responseUsers := responses.UserResponse{
			ID:    users.ID,
			Name:  users.Name,
			Email: users.Email,
		}
		usersList = append(usersList, responseUsers)
	}

	return usersList, nil
}

func (u *UserUseCase) GetUserByID(ctx context.Context, id uint) (responses.UserResponse, error) {
	log := logrus.WithContext(ctx)
	log.Info("Get user usecase")

	result, err := u.repository.FindByID(ctx, id)
	if err != nil {
		log.Error(err)
		return responses.UserResponse{}, err
	}

	log.Info("User found")
	responseUsers := responses.UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}

	return responseUsers, nil
}

func (u *UserUseCase) CreateUser(ctx context.Context, user requests.RegisterUserRequest) error {
	log := logrus.WithContext(ctx)
	log.Info("Create user usecase")

	userModel := models.UserModel{
		Name:  user.Name,
		Email: user.Email,
	}

	err := u.repository.Create(ctx, &userModel)

	if err != nil {
		log.WithError(err).Error("Error creating user")
		return err
	}

	log.Info("User created")
	return nil
}

func (u *UserUseCase) UpdateUser(ctx context.Context, id uint, user requests.UpdateUserRequest) error {
	log := logrus.WithContext(ctx)
	log.Info("Update user usecase")

	userModel := models.UserModel{
		Name:  user.Name,
		Email: user.Email,
	}

	err := u.repository.Update(ctx, &userModel)

	if err != nil {
		log.WithError(err).Error("Error updating user")
		return err
	}

	log.Info("User updated")
	return nil
}
