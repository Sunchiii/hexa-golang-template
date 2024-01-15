package handlers

import (
	"hexa-gorm/core/service"
	"hexa-gorm/packages/errs"
	"net/http"
  "strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.User
}

func NewUserHandler(userService service.User) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) GetUsers(c *gin.Context) {
	// Get all users
	users, err := h.userService.GetUsers()
	if err != nil {
    appError, ok := err.(*errs.AppError)
    if ok {
      c.JSON(appError.Code, appError)
      return
    }
    c.JSON(http.StatusInternalServerError, err)
		return
	}

  c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *userHandler) GetUser(c *gin.Context) {
  // Get user by id
  id := c.Param("id")
  // convert string to int
  idInt, err := strconv.Atoi(id)
  if err != nil {}
  user, err := h.userService.GetUser(idInt)
  if err != nil {
    appError, ok := err.(*errs.AppError)
    if ok {
      c.JSON(appError.Code, appError)
      return
    }
    c.JSON(http.StatusInternalServerError, err)
    return
  }
  c.JSON(http.StatusOK, gin.H{"data": user})
}


func (h *userHandler) CreateUser(c *gin.Context) {
  // Create user
  var userRequest service.UserRequest
  if err := c.ShouldBindJSON(&userRequest); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  user, err := h.userService.CreateUser(userRequest)
  if err != nil {
    appError, ok := err.(*errs.AppError)
    if ok {
      c.JSON(appError.Code, appError)
      return
    }
    c.JSON(http.StatusInternalServerError, err)
    return
  }
  c.JSON(http.StatusOK, gin.H{"data": user})
}
