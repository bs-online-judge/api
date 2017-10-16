package models

import "errors"

type User struct {
  Id        int
  Group     int8    `sql:",notnull,default:0"`
  Username  string  `sql:",notnull,unique"`
  Password  string  `sql:",notnull"`
}

func GetUserByCredentials(username, password string) (*User, error) {
  user := &User{}
  err := db.Model(user).Where("username = ?", username).Select()
  if err != nil {
    return nil, err
  }
  if user.Password != password {
    return nil, errors.New("invalid username/password")
  }
  return user, nil
}
