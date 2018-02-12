package main

import (

   "fmt"
   "encoding/json"
)

type person struct{
  Fname string
  Lname  string
  Age   int
  Rollno int
  option  []struct{
    Fname string
    Lname  string
    Age   int
    Rollno int
}

func main(){

  p1 := person{"prince","antony",21,35}
  bs,_ := json.Marshal(p1)
  fmt.Println("%T \n",bs)
  fmt.Println(string(bs))
}
