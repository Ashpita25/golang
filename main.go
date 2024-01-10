package main

import (
      "fmt"
      "database/sql"
      _ "github.com/go-sql-driver/mysql"
    )

func main () {

  db, err := sql.Open ("mysql", "root:@tcp(127.0.0.1:3306)/golang")
  if err != nil {
    panic(err)
  }

  defer db.Close ()
  fmt.Println ("Подключено к MySQL")
}
