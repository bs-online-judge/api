package models

import (
  "golang.org/x/crypto/bcrypt"
)

// user groups
const (
  RegisteredGroup = 0 // User's Group field default value
  WorkerGroup     = 1
  AdminGroup      = 2
)

// this table is read-heavy!
type User struct {
  Id            int
  Group         int8    `sql:",notnull,default:0"` // RegisteredGroup
  Username      string  `sql:",notnull,unique,type:varchar(255)"`
  PasswordHash  []byte  `sql:",notnull"`
}

func GetUserByCredentials(username, password string) (*User, error) {
  user := &User{}
  if err := db.Model(user).Where("username = ?", username).Select(); err != nil {
    return nil, err
  }
  if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
    return nil, err
  }
  return user, nil
}

func CreateUser(username, password string) error {
  passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return err
  }
  if err := db.Insert(&User{ Username: username, PasswordHash: passwordHash }); err != nil {
    return err
  }
  return nil
}
