package repository

import (
	"context"
	"errors"
	"user/models"

	"gorm.io/gorm"
)

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.Database.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByUserId(ctx context.Context, userId int64) (*models.User, error) {
	var user models.User

	err := r.Database.WithContext(ctx).Where("id = ?", userId).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) (int64, error) {
	err := r.Database.WithContext(ctx).Create(user).Error

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
