package models

import (
  "os"

  "github.com/go-pg/pg"
  "github.com/go-pg/pg/orm"
)

// db handle
var db *pg.DB

func CreateTables() error {
  db = pg.Connect(&pg.Options{
    Addr: os.Getenv("PG_HOSTNAME"),
    User: os.Getenv("PG_USERNAME"),
    Password: os.Getenv("PG_PASSWORD"),
    Database: "bsoj",
  })
  tables := []interface{} {
    &User{},
    &Problem{},
    &Submission{},
    &Task{},
    &Contest{},
  }
  for _, table := range tables {
    if err := db.CreateTable(table, &orm.CreateTableOptions{ IfNotExists: true }); err != nil {
      return err
    }
  }
  return nil
}
