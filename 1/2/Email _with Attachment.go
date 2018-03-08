package main

import (


  "gopkg.in/gomail.v2"


  )
  func main() {
    m := gomail.NewMessage()
    m.SetHeader("From", "princemcet@gmail.com")
    m.SetHeader("To", "princeantony70@gmail.com", "ravicse28@gmail.com","kchandruk7@gmail.com")
    m.SetAddressHeader("Cc", "princeantony70@outlook.com", "prince")
    m.SetHeader("Subject", "Hello!")
    m.SetBody("text/html", "Hello<h2> <b>chandru</b> and <i>Ravi</i></h2>!")
    m.Attach("/home/prince/Desktop/download.jpeg")

    d := gomail.NewPlainDialer("smtp.gmail.com", 587, "princemcet", "14bit035")

    // Send the email to.
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

}
