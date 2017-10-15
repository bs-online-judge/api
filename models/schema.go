package models

import (
  "os"
  "strings"
  
  "github.com/go-pg/pg"
)

// db handle
var db *pg.DB

func CreateTables() error {
  db = pg.Connect(&pg.Options{
    Addr: os.Getenv("PG_ADDR"),
    User: os.Getenv("PG_USER"),
    Password: os.Getenv("PG_PASSWORD"),
    Database: "bsoj",
  })
  tables := []interface{} {
    &Task{},
  }
  for _, table := range tables {
    err := db.CreateTable(table, nil)
    if err != nil && !strings.Contains(err.Error(), "already exists") {
      return err
    }
  }
  return nil
}
