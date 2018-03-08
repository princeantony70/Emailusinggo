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


       type Database struct {

         Username  string  `json: "username"`
         Fbid     string  `json: "name"`

      }

      type Getbase struct{
       Id    int         `json: "id"`
			 Username  string  `json: "username"`
			 Fbid     string   `json: "name"`
			 Is_active int     `json: "is_active"`

			}


      type userAddHandler struct {
      	db *sql.DB
      }
			type userGetHandler struct {
				db *sql.DB
			}

var  err error

func (u userAddHandler) insertInDatabase(data Database) error {
	   _, err := u.db.Exec("INSERT INTO persons(username,fbid) VALUES(?, ?)", data.Username, data.Fbid)
if err != nil {
		return err
}
fmt.Println(data.Username)
return err
}

       func (u userAddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
       dec := json.NewDecoder(r.Body)
       defer r.Body.Close()

       var k Database
       if err := dec.Decode (&k); err != nil {
       	fmt.Println("unmarshall error ", err)
       }

			 if err := u.insertInDatabase(k); err != nil {
			 	fmt.Fprintln(w, err.Error())
			 return
			 }

       fmt.Fprintln(w, `{"code ":"success"}`)
       }



func (v userGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

var get []Getbase

rows, err := v.db.Query("select id, username, fbid,is_active from persons;")
		if err != nil {
			fmt.Print(err.Error())
		}

		for rows.Next() {
			var id ,is_active int
			var username,fbid string


			rows.Scan(&id ,&username, &fbid, &is_active)
			get = append(get, Getbase{id, username, fbid, is_active})

			if err != nil {
				fmt.Print(err.Error())
			}
		}
		usersBytes, _ := json.Marshal(&get)

		w.Write(usersBytes)

	 }







       func main() {
       	db, err := sql.Open("mysql", "root:nfn@tcp(127.0.0.1:3306)/auto")
       	if err != nil {
       		log.Fatalf("failed to open db: %s", err)
       	}
       	err = db.Ping()
       	if err != nil {
       		log.Fatal(err)
       	}
       	defer db.Close()

       	handler := userAddHandler{
       		db: db,
       	}
				handler2 := userGetHandler{
					db: db,
				}
       	http.Handle("/add", handler)
				http.Handle("/get",handler2)
       	http.ListenAndServe(":6066", nil)
       }
