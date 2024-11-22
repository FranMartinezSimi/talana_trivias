package repository

import (
	"context"
	"talana_prueba_tecnica/src/entity/models"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type UserRepository struct {
	gorm *gorm.DB
}

func NewUserRepository(gorm *gorm.DB) *UserRepository {
	return &UserRepository{gorm}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]models.UserModel, error) {
	log.WithContext(ctx).Println("finding all users")
	var users []models.UserModel

	log.Info("finding all users")

	res := r.gorm.WithContext(ctx).Find(&users)
	if res.Error != nil {
		log.Error("Error finding all users")
		return nil, res.Error
	}

	log.WithError(res.Error).Info("users found")
	return users, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uint) (*models.UserModel, error) {
	log.WithContext(ctx).Println("finding user by id")
	var user models.UserModel

	log.Info("finding user by id")

	res := r.gorm.WithContext(ctx).First(&user, id)
	if res.Error != nil {
		log.Error("Error finding user by id")
		return nil, res.Error
	}

	log.WithError(res.Error).Info("user found")
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.UserModel) error {
	log.WithContext(ctx).Println("creating user")

	log.Info("creating user")

	res := r.gorm.WithContext(ctx).Create(user)
	if res.Error != nil {
		log.Error("Error creating user")
		return res.Error
	}

	log.WithError(res.Error).Info("user created")
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.UserModel) error {
	log.WithContext(ctx).Println("updating user")

	log.Info("updating user")

	res := r.gorm.WithContext(ctx).Save(user)
	if res.Error != nil {
		log.Error("Error updating user")
		return res.Error
	}

	log.WithError(res.Error).Info("user updated")
	return nil
}
