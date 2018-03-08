package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "io/ioutil"
	"log"
	"net/http"
)

type Tag struct {
	ID int `json:"id"`
}

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


// type Response struct {
// 	Status  bool        `json:"status"`
// 	Types   interface{}  `json:"types"`
// 	Formula interface{} `json:"formula"` //Replace this with your actual type
// }
// */
type Response struct {
	Status bool `json:"status"`
	Value        `json:"value"`
}

type	Value  struct {
		Types   []string `json:"types"`
		Formula []string `json:"formula"`
	}


type Validation struct{
	Format string `json:"format"`
	Regx   string  `json:"regx"`
}


type userAddHandler struct {
	db *sql.DB
}

type userGetHandler struct {
	db *sql.DB
}

func (u userAddHandler) insertInDatabase(data Questions) error {

results, err := u.db.Exec("INSERT INTO profile_questions(name, section, position,title,titleSpanish,submited_value,spanish_submited_value,des,ans,view_type,parent_id,isRequired,is_submit_field,is_active) VALUES(?, ?, ?,?,?,?,?,?,?,?,?,?,?,?)", data.Question.Name, data.Question.Section, data.Question.Position, data.Question.Title, data.Question.TitleSpanish, data.Question.SubmitedValue, data.Question.SpanishSubmitedValue, data.Question.Des, data.Question.Ans, data.Question.ViewType, data.Question.ParentID, data.Question.IsRequired, data.Question.IsSubmitField, data.Question.IsActive)
if err != nil {
	return err
}
if len(data.Options) > 0 {

lastInsertId , err := results.LastInsertId()
if err != nil {
return err
}

for _, option := range data.Options {
		_, err = u.db.Exec("INSERT INTO profile_questions(name, section, position,title,titleSpanish,submited_value,spanish_submited_value,des,ans,view_type,parent_id,isRequired,is_submit_field,is_active) VALUES(?, ?, ?,?,?,?,?,?,?,?,?,?,?,?)",
				   option.Name,
				   option.Section,
				   option.Position,
				   option.Title,
				   option.TitleSpanish,
				   option.SubmitedValue,
				   option.SpanishSubmitedValue,
				   option.Des,
				   option.Ans,
				   option.ViewType,
				   lastInsertId,
				   option.IsRequired,
				   option.IsSubmitField,
				   option.IsActive)
if err != nil {
		return err
}
}
}else if data.Validation.Messgae != "" {
		_, err = u.db.Exec("INSERT INTO input_types(messgae,messageSpanish,regx,format) VALUES(?,?,?,?)", data.Validation.Messgae, data.Validation.MessageSpanish, data.Validation.Regx, data.Validation.Format)
	}

	return err
}

func (u userAddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
dec := json.NewDecoder(r.Body)
defer r.Body.Close()

var k Questions
if err := dec.Decode (&k); err != nil {
	fmt.Println("unmarshall error ", err)
}

if err := u.insertInDatabase(k); err != nil {
	fmt.Fprintln(w, err.Error())
return
}
	// no need for extra casting

fmt.Fprintln(w, `{"code ":"success"}`)
}



func (v userGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {




rows, err := v.db.Query("SELECT  DISTINCT format  FROM input_types")
if err != nil {
panic(err)
}
types := make([]string, 0, 8) // This creates a slice of strings with a capacity of 8 and length of 0.
defer rows.Close()
for rows.Next() {
validation := Validation{}
err = rows.Scan(&validation.Format)
if err != nil {
fmt.Println("failed to scan validation data", err)
}
types = append(types, validation.Format) // This appends the formula to the slice of formulas.
//fmt.Println(types)
}

rf, err := v.db.Query("SELECT  DISTINCT regx  FROM input_types")
if err != nil {
panic(err)
}
formulas := make([]string, 0, 8) // This creates a slice of strings with a capacity of 8 and length of 0.
defer rows.Close()
for rf.Next() {
validation := Validation{}
err = rf.Scan(&validation.Regx)
if err != nil {
	fmt.Println("failed to scan validation data ", err)
}
formulas = append(formulas, validation.Regx) // This appends the formula to the slice of formulas.
//fmt.Println(formulas)
}

	//values := validation.Format   //Do some SQL queries to get response here
	//formulas := validation.Regx// Do some SQL queries to get response here
	response := Response{
		Status:  true,
		Value :Value{
		Types:   types ,
		Formula: formulas,
	},
}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
     }



func main() {
	db, err := sql.Open("mysql", "root:nfn@tcp(127.0.0.1:3306)/shift_pixy")
	if err != nil {
		log.Fatalf("failed to open db: %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		//Ping is needed to test the conection - Open only tests the connection string
	}
	defer db.Close()

	handler := userAddHandler{
		db: db,
	}
	handler2 := userGetHandler{
		db: db,
	}

	http.Handle("/add", handler)
	http.Handle("/get", handler2)
	http.ListenAndServe(":1351", nil)
}
