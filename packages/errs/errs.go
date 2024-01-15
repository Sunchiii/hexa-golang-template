package errs

import "net/http"

type AppError struct {
  Code int
  Message string
}

func (a AppError) Error() string {
  return a.Message
}

func NewNotFoundError(message string) error {
  return AppError{
    Code: http.StatusNotFound,
    Message: message,
  }
}

func NewUnexpectedError() error {
  return AppError{
    Code: http.StatusInternalServerError,
    Message: "unexpected error",
  }
}

func NewValidationError(message string) error {
  return AppError{
    Code: http.StatusUnprocessableEntity,
    Message: message,
  }
}

func NewAuthenticationError(message string) error {
  return AppError{
    Code: http.StatusUnauthorized,
    Message: message,
  }
}

func NewBadRequestError(message string) error {
  return AppError{
    Code: http.StatusBadRequest,
    Message: message,
  }
}

func NewForbiddenError(message string) error {
  return AppError{
    Code: http.StatusForbidden,
    Message: message,
  }
}
