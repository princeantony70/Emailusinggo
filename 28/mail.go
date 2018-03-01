package main


import (
	"log"
	"net/smtp"
  "encoding/json"
)

func main() {
http.Handle("/send",send)
send("hello there")
}




// type  Request struct{
// 	From    string  `json:"from"`
// 	Password string `json:"password"`
// 	To      string  `json:"to"`
// 	Message string  `json:"message"`
// }


func send(body string) {

  from := "princemcet@gmail.com"
	pass := "14bit035"
	to := "ravicse28@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
