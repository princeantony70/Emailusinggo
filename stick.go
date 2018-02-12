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
    Parent_id            int    `json:"parent_id"`
		Name                 string `json:"name"`
		Section              string `json:"section"`
		Rollno               int    `json:"rollno"`
		Age                  int    `json:"age"`
	} `json:"question"`
	Options []struct {
    Parent_id            int    `json:"parent_id"`
		Name                 string `json:"name"`
		Section              string `json:"section"`
		Rollno               int    `json:"rollno"`
		Age                  int    `json:"age "`

	} `json:"options"`
}



func insertInDatabase(data Questions) error  {

  for i:=0;i<=1;i++{

     switch i {
     case 0:
       _, err = appdatabase.Exec("INSERT INTO array(parent_id,name, section, rollno,age) VALUES(?,?,?,?,?)",data.Question.Parent_id,data.Question.Name,data.Question.Section , data.Question.Rollno,data.Question.Age)

      case 1:

      for i:=0;i<=1;i++{
      _, err = appdatabase.Exec("INSERT INTO array(parent_id,name, section, rollno,age) VALUES(?,?,?,?,?)",data.Options[i].Parent_id,data.Options[i].Name ,data.Options[i].Section,data.Options[i].Rollno,data.Options[i].Age)
   }
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
}


func init(){
appdatabase, err = sql.Open("mysql", "root:nfn@/json")
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
       http.ListenAndServe(":6023", nil)
}
