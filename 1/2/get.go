package main

import "log"
import "net/http"
import "encoding/json"

import (
    _"github.com/go-sql-driver/mysql"
    "database/sql"
)


func init(){
db, err := sql.Open("mysql", "root:nfn@(127.0.0.1:3306)/auto")

}

func questions(w http.ResponseWriter, r *http.Request) {
    // db, err := sql.Open("mysql", "<username>:<password>@tcp(127.0.0.1:<port>)/<dbname>?charset=utf8" )

    w.Header().Set("Content-Type", "application/json")
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("select id, name, section , position, title, submitted_value,des, ans,view_type,parent_id,is_required,is_submit_field,is_active,created_at,updated_at from questions")

    if err != nil {
        log.Fatal(err)
    }

    type Questions struct {

      ID int  `json: "id"`
      Name string  `json: "name"`
      Section string  `json: "section"`
      Position int `json: "position"`
      Title string  `json: "title"`
      Submitted_Value string `json: "submitted_value"`
      Des string `json: "des"`
      Ans string  `json: "ans"`
      View_type int `json: "view_type"`
      Parent_id int  `json: "parent_id"`
      Is_required int `json: "is_required"`
      Is_submit_field int `json: "is_submit_field"`
      Is_active int  `json: "is_active"`

   }

   var qus []Questions

       for rows.Next() {
           var id, position, view_type, parent_id, is_required, is_submit_field, is_active int
           var name string
           var section, title,submitted_value string
           var des,ans string

           rows.Scan(&id ,&name, &section, &position, &title, &submitted_value, &des,&ans,&view_type,&parent_id,&is_required,&is_submit_field,&is_active)
           qus = append(qus, Questions{id, name, section, position, title, submitted_value,des,ans,view_type,parent_id,is_required,is_submit_field,is_active})
       }

       usersBytes, _ := json.Marshal(&qus)

       w.Write(usersBytes)
       db.Close()
      }



      func main() {
       http.HandleFunc("/questions/", questions)
       http.ListenAndServe(":8082", nil)
       }
