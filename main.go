package main

import "fmt"

func main(){
  v := 5
  p := &v
  
  fmt.Print(*p, " ")
  changePointer(p)
  fmt.Println(v)
}
func changePointer(j *int){
  v := 3
  j = &v
}