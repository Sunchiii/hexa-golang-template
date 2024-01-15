package repository

import "hexa-gorm/core/models"

type UserRepo interface{
  GetUsers() ([]models.User, error)
  GetUser(id int) (*models.User, error)
  CreateUser(user models.User) (*models.User, error)
}

