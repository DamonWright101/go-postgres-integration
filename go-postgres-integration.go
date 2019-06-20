package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "MySecureUser"
  password = "AVerySecureassword"
  dbname   = "app"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)

  // there be errors
  if err != nil {
    panic("params are syntaxly wrong == \n" + err)
  }
  defer db.Close()

  err = db.Ping()

  if err != nil {
    panic("Server is not alive == \n" + err)
  }

  fmt.Println("Got connected!")
}