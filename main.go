package main

import ("fmt"
    "net/http"
    "html/template"
    "database/sql"
     _ "github.com/go-sql-driver/mysql")

func index(w http.ResponseWriter, r *http.Request){
        t, err :=  template.ParseFiles ("templates/index.html", "templates/header.html", "templates/footer.html")

        if err != nil {
        fmt.Fprintf (w, err.Error())
        }
         t.ExecuteTemplate(w, "index", nil)
}
func create(w http.ResponseWriter, r *http.Request){
        t, err :=  template.ParseFiles ("templates/create.html", "templates/header.html", "templates/footer.html")

        if err != nil {
        fmt.Fprintf (w, err.Error())
        }
         t.ExecuteTemplate(w, "create", nil)
}
func save_user(w http.ResponseWriter, r *http.Request){
        user := r.FormValue("name")
        pass := r.FormValue("pass")
        age := r.FormValue("age")
        db, err := sql.Open ("mysql", "root:@tcp(127.0.0.1:3306)/golang") //подключение к БД
          if err != nil {
            panic(err)
          }
        defer db.Close ()
          //Добавление данных
          insert, err := db.Query (fmt.Sprintf("INSERT INTO `userDB`(`name`,`pass`, `age`) VALUES ('%s','%s','%s')", user, pass, age))
          if err != nil{
            panic (err)
            }
            defer insert.Close()
            http.Redirect (w,r,"/", http.StatusSeeOther)

}

func handleRequest() {
    http.HandleFunc("/", index)
    http.HandleFunc ("/create", create)
    http.HandleFunc ("/save_user", save_user)
    http.ListenAndServe(":8080", nil)
}
func main () {
    //bob := User {name: "Bob", age:25, money:-50, avg_grades: 4.2, hapiness: 0.8}

    handleRequest()
}