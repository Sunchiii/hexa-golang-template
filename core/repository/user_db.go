package repository

import (
	"hexa-gorm/core/models"

	"gorm.io/gorm"
)

type userRepository struct {
  db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepo {
  return userRepository{db}
}

func (r userRepository) GetUsers() ([]models.User, error) {
  var users []models.User
  err := r.db.Find(&users).Error
  if err != nil {
    return nil, err
  }
  return users, nil
}

func (r userRepository) GetUser(id int) (*models.User, error) {
  var user models.User
  err := r.db.First(&user, id).Error
  if err != nil {
    return nil, err
  }
  return &user, nil
}

func (r userRepository) CreateUser(user models.User) (*models.User, error) {
  err := r.db.Create(&user).Error
  if err != nil {
    return nil, err
  }
  return &user, nil
}
