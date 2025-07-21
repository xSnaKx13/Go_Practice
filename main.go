package main

import (
	"fmt"
)

func main(){
  var a, b int
  _, err := fmt.Scan(&a, &b)
  if err != nil{
    fmt.Println("ошибка")
    return
  }

  res, err := devide(a, b)
  fmt.Println(res)
}

func devide(a,b int) (int, error){
  return a/b, nil
}