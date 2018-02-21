package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
_ "github.com/go-sql-driver/mysql"
)


var appdatabase  *sql.DB
var err error
var arr int


type Tag struct {
  ID   int    `json:"id"`

}


type Questions struct {
	Question struct {
		Name                 string `json:"name"`
		Section              string      `json:"section"`
		Position             int         `json:"position"`
		Title                string      `json:"title"`
		TitleSpanish         string      `json:"titleSpanish"`
		SubmitedValue        string      `json:"submited_value"`
		SpanishSubmitedValue string      `json:"spanish_submited_value"`
		Des                  string      `json:"des"`
		Ans                  string      `json:"ans"`
		ViewType             int         `json:"view_type"`
		ParentID             int         `json:"parent_id"`
		IsRequired           int         `json:"isRequired"`
		IsSubmitField        int         `json:"is_submit_field"`
		IsActive             int         `json:"is_active"`
	} `json:"question"`
	Options []struct {
		Name                 string `json:"name"`
		Section              string      `json:"section"`
		Position             int         `json:"position"`
		Title                string      `json:"title"`
		TitleSpanish         string      `json:"titleSpanish"`
		SubmitedValue        string      `json:"submited_value"`
		SpanishSubmitedValue string      `json:"spanish_submited_value"`
		Des                  string      `json:"des"`
		Ans                  string      `json:"ans"`
		ViewType             int         `json:"view_type"`
		IsRequired           int         `json:"isRequired"`
		IsSubmitField        int         `json:"is_submit_field"`
		IsActive             int         `json:"is_active"`
	} `json:"options"`
	Validation struct {
		Messgae        string `json:"messgae"`
		MessageSpanish string `json:"messageSpanish"`
		Regx           string `json:"regx"`
		Format         string `json:"format"`
	} `json:"validation"`
}

func insertInDatabase(data Questions) error {

if len(data.Options) > 0{

_, err = appdatabase.Exec("INSERT INTO profile_questions(name, section, position,title,titleSpanish,submited_value,spanish_submited_value,des,ans,view_type,parent_id,isRequired,is_submit_field,is_active) VALUES(?, ?, ?,?,?,?,?,?,?,?,?,?,?,?)", data.Question.Name ,data.Question.Section ,data.Question.Position,data.Question.Title,data.Question.TitleSpanish,data.Question.SubmitedValue,data.Question.SpanishSubmitedValue,data.Question.Des,data.Question.Ans,data.Question.ViewType,data.Question.ParentID,data.Question.IsRequired,data.Question.IsSubmitField,data.Question.IsActive)
}

return err
}

func userAddHandler(w http.ResponseWriter, r *http.Request) {
body, err := ioutil.ReadAll(r.Body)
  if err != nil {
	fmt.Println("reading the Request error")
    }
var k Questions
  err = json.Unmarshal(body, &k)
  if err != nil {
      fmt.Println("unmarshall error ")
  }
	err = insertInDatabase(k)
	w.Write([]byte(`{"code ":"success"}`))
}


func init() {
fmt.Println("Go MySQL connection status :")
appdatabase, err := sql.Open("mysql", "root:nfn@tcp(127.0.0.1:3306)/api")
if err != nil {
fmt.Println("db connection error")
    }
			    defer appdatabase.Close()
}

func main() {
	http.HandleFunc("/add", userAddHandler)
	http.ListenAndServe(":1221",nil)
}
