package service

// for response
type UserResponse struct {
  ID uint `json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
  Picture string `json:"picture"`
}

// for request
type UserRequest struct {
  Name string `json:"name" binding:"required"`
  Email string `json:"email" binding:"required"`
  Password string `json:"password" binding:"required"`
  Picture string `json:"picture"`
}

type User interface {
  GetUsers() ([]UserResponse, error)
  GetUser(id int) (*UserResponse, error)
  CreateUser(user UserRequest) (*UserResponse, error)
}
