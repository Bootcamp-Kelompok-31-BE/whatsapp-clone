package repositories

import (
	"context"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/messaging/models"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/util"
)

type IUserRepository interface {
	GetUserByID(ctx context.Context, userID uint) (*models.User, error)
}

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID uint) (*models.User, error) {
	db := util.GetDBFromContext(ctx)

	var user models.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
