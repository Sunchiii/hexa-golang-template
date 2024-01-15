package service

import (
	"database/sql"
	"hexa-gorm/core/models"
	"hexa-gorm/core/repository"
	"hexa-gorm/packages/errs"
	"hexa-gorm/packages/logs"
)

type userService struct {
  repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) User {
  return userService{repo}
}

func (s userService) GetUsers() ([]UserResponse, error) {
  var users []UserResponse

  result, err := s.repo.GetUsers()
  if err != nil {
    if err == sql.ErrNoRows{
      return []UserResponse{}, nil
    }
    logs.Error(err)
    return nil, err
  }

  // convert models to dto
  for _, user := range result {
    users = append(users, UserResponse{
      ID: user.ID,
      Name: user.Name,
      Email: user.Email,
      Picture: user.Picture,
    })
  }
  return users, nil
}

func (s userService) GetUser(id int) (*UserResponse, error) {
  user, err := s.repo.GetUser(id)
  if err != nil {
    if err == sql.ErrNoRows{
      return nil, errs.NewNotFoundError("user not found")
    }
    logs.Error(err)
    return nil, err
  }

  // convert model to dto
  return &UserResponse{
    ID: user.ID,
    Name: user.Name,
    Email: user.Email,
    Picture: user.Picture,
  }, nil
}

func (s userService) CreateUser(request UserRequest) (*UserResponse, error) {
  // convert dto to model
  user := models.User{
    Name: request.Name,
    Email: request.Email,
    Password: request.Password,
    Picture: request.Picture,
  }
  result, err := s.repo.CreateUser(user)
  if err != nil {
    logs.Error(err)
    return nil, err
  }
  // convert model to dto
  return &UserResponse{
    ID: result.ID,
    Name: result.Name,
    Email: result.Email,
    Picture: result.Picture,
  }, nil
}
