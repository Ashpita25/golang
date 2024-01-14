package main

import (
      "fmt"
      //"database/sql"
      "net/http"
      "html/template"
      //_ "github.com/go-sql-driver/mysql"
    )

    func create (w http.ResponseWriter, r *http.Request) {
            t, err := template.ParseFiles ("templates/index.html", "templates/header.html", "templates/footer.html")
            if err!= nil {
            fmt.Println (w, err.Error())
            }
            t.ExecuteTemplate (w, "index", nil)
    }

func index (w http.ResponseWriter, r *http.Request) {
        t, err := template.ParseFiles ("templates/index.html", "templates/header.html", "templates/footer.html")
        if err!= nil {
        fmt.Println (w, err.Error())
        }
        t.ExecuteTemplate (w, "index", nil)
}
func handleFunc () {
    http.HandleFunc ("/", index)
    http.HandleFunc ("/create", create)
    http.ListenAndServe (":8080", nil)
}
func main () {
        handleFunc ()
}

