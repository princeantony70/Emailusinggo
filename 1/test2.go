package main

import (
    "fmt"
    "log"
    "net/http"
    "net/smtp"
    "encoding/json"
)

type Request struct{
   From string     `json:"from"`
   Password string `json:"password"`
   To  string      `json:"to"`
   Message string ` json:"message"`
}


/*  type Response struct {
  	 Status bool `json:"status"`
  	 Code   string    `json:"code"`
}*/


func getrequest(w http.ResponseWriter, r *http.Request){

    dec := json.NewDecoder(r.Body)
    defer r.Body.Close()

   var content Request
   if err := dec.Decode (&content); err != nil {
  	fmt.Println("unmarshall error ", err)
   }


   msg := "From: " + content.From + "\n" +
		"To: " + content.To + "\n" +
		"Subject: Hello there\n\n" +
		content.Message


    /*var code1  string
    var code2 string

    code1 = "mail successfully send"
    code2 = "invalid feilds"

    response1 := Response{
      Status :  true,
      Code : code1,
    }

    response2 := Response{
      Status :  false,
      Code : code2,
    }
    w.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(w)
    //encoder.Encode(response)*/

	  err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", content.From, content.Password, "smtp.gmail.com"),
		content.From, []string{content.To}, []byte(msg))


	if err != nil {
    w.Write([]byte(`{"Status":"false","code":"Invalid field"}`))
    // w.Write([]byte(`,"code":"Invalid feild"}`))


    //encoder.Encode(response2)
    //fmt.Fprintf(w, "smptp error %s!", err)
    //log.Printf("smtp error: %s", err)
		return
	}
  w.Write([]byte(`{"Status":"true","code":"mail send "}`))
  // w.Write([]byte(`{:`))

  }

func main() {
    http.HandleFunc("/send", getrequest)
    log.Fatal(http.ListenAndServe(":8093", nil))
}
