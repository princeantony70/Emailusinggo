package main

import (
      "encoding/json"
      "net/http"
      "fmt"
      "io"
     _ "github.com/go-sql-driver/mysql"
      "database/sql"


)

var appdatabase *sql.DB
var err error

type Questions struct {
	Question struct {
		Name                 interface{} `json:"name"`
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

	Validation struct {
		Messgae        string `json:"messgae"`
		MessageSpanish string `json:"messageSpanish"`
		Regx           string `json:"regx"`
		Format         string `json:"format"`
	} `json:"validation"`
}



func insertInDatabase(data Questions) error  {
 for i:=0;i<=1;i++{
     switch i {
     case 0:
_, err = appdatabase.Exec("INSERT   INTO input_types(name,section,position,title,titleSpanish,submitted_value,spanish_submited_value,des,ans,view_type,parent_id,is_required,is_submit_field,is_active)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)",data.Question.Name,data.Question.Section,data.Question.Position,data.Question.Title,data.Question.TitleSpanish,data.Question.SubmitedValue,data.Question.SpanishSubmitedValue,data.Question.Des,data.Question.Ans,data.Question.ViewType,data.Question.ParentID,data.Question.IsRequired,data.Question.IsSubmitField,data.Question.IsActive)

     case 1:
     _, err = appdatabase.Exec("INSERT INTO input_types(messgae,messageSpanish,regx,format) VALUES(?,?,?,?)",data.Validation.Messgae,data.Validation.MessageSpanish,data.Validation.Regx,data.Validation.Format)
}
}
return err
}

func userAddHandler(w http.ResponseWriter, r *http.Request) {


      //make byte array
      out := make([]byte,1024)
      bodyLen, err := r.Body.Read(out)

      if err != io.EOF {
             fmt.Println(err.Error())
             w.Write([]byte("{error:" + err.Error() + "}"))
             return
      }

      var k Questions

      err = json.Unmarshal(out[:bodyLen],&k)


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


func init(){
appdatabase, err = sql.Open("mysql", "root:nfn@/validation")
if err != nil{
 fmt.Println("db error ")
}
err = appdatabase.Ping()
if err !=nil{
 fmt.Println("ping error")
}
}


func main() {
      http.HandleFunc("/add", userAddHandler)
      http.ListenAndServe(":6033", nil)
}
