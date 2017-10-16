package models

import (
  "errors"
  
  //"golang.org/x/crypto/bcrypt"
)

// user groups
const (
  RegisteredGroup = 0 // User's Group field default value
  WorkerGroup     = 1
  AdminGroup      = 2
)

type User struct {
  Id            int
  Group         int8    `sql:",notnull,default:0"` // RegisteredGroup
  Username      string  `sql:",notnull,unique,type:varchar(255)"`
  Password  string  `sql:",notnull,type:varchar(255)"`
}

func GetUserByCredentials(username, password string) (*User, error) {
  user := &User{}
  if err := db.Model(user).Where("username = ?", username).Select(); err != nil {
    return nil, err
  }
  if user.Password != password {
    return nil, errors.New("invalid username/password")
  }
  return user, nil
}
