package models

import "gorm.io/gorm"

// for database
type User struct {
  gorm.Model
  Name string `db:"name"`
  Email string `db:"email"`
  Password string `db:"password"`
  Picture string `db:"picture"`
}


