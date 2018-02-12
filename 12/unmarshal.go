package main

import (
     "encoding/json"
     "fmt"
)



type Person struct {
  First  string
  Last   string
  Age    int
  Rollno  int
}


func main(){
  var p1 Person
  fmt.Println(p1.First)
  fmt.Println(p1.Last)
  fmt.Println(p1.Age)
  fmt.Println(p1.Rollno)
  bs := []byte(`{"First:name","Last":name","Age":21,"Rollno":35}`)
  json.Unmarshal(bs, &p1)




  
fmt.Println("----------------------------------------------------------")
fmt.Println(p1.First)
fmt.Println(p1.Last)
fmt.Println(p1.Age)
fmt.Println(p1.Rollno)
fmt.Printf("%T \n",p1)
}
