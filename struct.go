package main

import "fmt"

type person struct {
  fname string
  lname string
}

type doublezeroseven struct{
  person
  license bool
}

func (p person) fullname() string{

  return p.fname+p.lname

}

func ( p person) fullname() string  {



}

func main(){
 p1 := person{"prince","antony"}
 p2 :=person{"jeffry","rohan"}
 fmt.Println(p1.fullname())
fmt.Println(p2.fullname())

}
