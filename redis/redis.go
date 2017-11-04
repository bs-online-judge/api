package redis

import (
  "os"
  "time"
  "encoding/json"

  goredis "gopkg.in/redis.v5"

  "github.com/bs-online-judge/api/models"
)

// TODO: maybe we drop it
const (
  RedisNoUser = 100
)

type UserSession struct {
  UserId int `json:"userId"`
}

var db *goredis.Client

func Connect() error {
  db = goredis.NewClient(&goredis.Options{
    Addr: os.Getenv("REDIS_HOSTNAME"),
    Password: os.Getenv("REDIS_PASSWORD"),
    DB: 0,
  })

  _, err := db.Ping().Result()
  if err != nil {
    return err
  }
  return nil
}

func GetUserSession(key string) (*UserSession, error) {
  raw, err := db.Get(key).Result()
  if err != nil {
    return nil, err
  }

  user := &UserSession{}
  if err := json.Unmarshal([]byte(raw), user); err != nil {
    return nil, err
  }

  return user, nil
}

func SetUserSession(key string, user *models.User) error {
  raw, err := json.Marshal(&UserSession{user.Id})
  if err != nil {
    return err
  }

  if err := db.Set(key, string(raw), time.Hour).Err(); err != nil {
    return err
  }

  return nil
}
