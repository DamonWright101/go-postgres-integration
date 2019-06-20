package main

import (
  "os"
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    os.Getenv("POSTGRES_HOST"), 
    os.Getenv("POSTGRES_PORT"), 
    os.Getenv("POSTGRES_USER"), 
    os.Getenv("POSTGRES_PASSWORD"), 
    os.Getenv("POSTGRES_DB"))

  db, err := sql.Open("postgres", psqlInfo)

  // there be errors
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()

  if err != nil {
    panic(err)
  }

  fmt.Println("Got connected!")
}