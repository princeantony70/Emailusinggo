package main

import (
	"bytes"
	"html/template"
	"log"

	gomail "gopkg.in/gomail.v2"
)

type info struct {
	Name string
}

func (i info) sendMail() {

	t := template.New("template.html")

	var err error
	t, err = t.ParseFiles("template.html")
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, i); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "princemcet@gmail.com")
	m.SetHeader("To", "ravicse28@gmail.com","kchandruk7@gmail.com","princeantony70@gmail.com")
	m.SetAddressHeader("Cc", "princeantony70@outlook.com", "prince")
	m.SetHeader("Subject", "golang test")
	m.SetBody("text/html", result)
	m.Attach("template.html")// attach whatever you want

	d := gomail.NewDialer("smtp.gmail.com", 587, "princemcet@gmail.com", "14bit035")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
func main() {

	d := info{"prince"}

	d.sendMail()
}
