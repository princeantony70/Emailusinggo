package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var appdatabase *sql.DB
var err error
var arr int

type Questions struct {
	Question struct {
		Name                 string `json:"name"`
		Section              string `json:"section"`
		Position             int    `json:"position"`
		Title                string `json:"title"`
		TitleSpanish         string `json:"titleSpanish"`
		SubmitedValue        string `json:"submited_value"`
		SpanishSubmitedValue string `json:"spanish_submited_value"`
		Des                  string `json:"des"`
		Ans                  string `json:"ans"`
		ViewType             int    `json:"view_type"`
		ParentID             int    `json:"parent_id"`
		IsRequired           int    `json:"isRequired"`
		IsSubmitField        int    `json:"is_submit_field"`
		IsActive             int    `json:"is_active"`
	} `json:"question"`
	Options []struct {
		Name                 string `json:"name"`
		Section              string `json:"section"`
		Position             int    `json:"position"`
		Title                string `json:"title"`
		TitleSpanish         string `json:"titleSpanish"`
		SubmitedValue        string `json:"submited_value"`
		SpanishSubmitedValue string `json:"spanish_submited_value"`
		Des                  string `json:"des"`
		Ans                  string `json:"ans"`
		ViewType             int    `json:"view_type"`
		IsRequired           int    `json:"isRequired"`
		IsSubmitField        int    `json:"is_submit_field"`
		IsActive             int    `json:"is_active"`
	} `json:"options"`
	Validation struct {
		Messgae        string `json:"messgae"`
		MessageSpanish string `json:"messageSpanish"`
		Regx           string `json:"regx"`
		Format         string `json:"format"`
	} `json:"validation"`
}

type Tag struct {
	ID int `json:"id"`
}

func insertInDatabase(data Questions) error {

	if len(data.Options) > 0 {
		fmt.Println(data.Options)
	} else {
		fmt.Println(data.Question)
	}

	for i := 0; i <= 1; i++ {
		switch i {
		case 0:
			_, err = appdatabase.Exec("INSERT INTO profile_questions(name, section, position,title,titleSpanish,submited_value,spanish_submited_value,des,ans,view_type,parent_id,isRequired,is_submit_field,is_active) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)", data.Question.Name, data.Question.Section, data.Question.Position, data.Question.Title, data.Question.TitleSpanish, data.Question.SubmitedValue, data.Question.SpanishSubmitedValue, data.Question.Des, data.Question.Ans, data.Question.ViewType, data.Question.ParentID, data.Question.IsRequired, data.Question.IsSubmitField, data.Question.IsActive)
			results, err := appdatabase.Query("SELECT LAST_INSERT_ID()")
			if err != nil {
				panic(err.Error())
			}

			for results.Next() {
				var tag Tag
				err = results.Scan(&tag.ID)
				if err != nil {
					panic(err.Error())
				}
				arr = tag.ID
				fmt.Println(arr)
			}

		case 1:
			for i := 0; i <= 1; i++ {
				_, err = appdatabase.Exec("INSERT INTO profile_questions(name, section, position,title,titleSpanish,submited_value,spanish_submited_value,des,ans,view_type,parent_id,isRequired,is_submit_field,is_active) VALUES(?, ?, ?,?,?,?,?,?,?,?,?,?,?,?)", data.Options[i].Name, data.Options[i].Section, data.Options[i].Position, data.Options[i].Title, data.Options[i].TitleSpanish, data.Options[i].SubmitedValue, data.Options[i].SpanishSubmitedValue, data.Options[i].Des, data.Options[i].Ans, data.Options[i].ViewType, arr, data.Options[i].IsRequired, data.Options[i].IsSubmitField, data.Options[i].IsActive)

			}
		}
	}

	return err
}

func userAddHandler(w http.ResponseWriter, r *http.Request) {

	//make byte array
	out := make([]byte, 1024)
	bodyLen, err := r.Body.Read(out)

	if err != io.EOF {
		fmt.Println(err.Error())
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	var k Questions

	err = json.Unmarshal(out[:bodyLen], &k)

	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	err = insertInDatabase(k)

	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	w.Write([]byte(`{"code ":"success"}`))
	fmt.Println(k)

}

func init() {
	appdatabase, err = sql.Open("mysql", "root:nfn@/api")
}

func main() {
	http.HandleFunc("/add", userAddHandler)
	http.ListenAndServe(":7082", nil)
}
