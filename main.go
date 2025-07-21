package main

import (
	"errors"
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
  if res != 0{
    fmt.Println(res)
  }else{
    fmt.Println(err)
  }
}

func devide(a,b int) (int, error){
  if a == 0 || b == 0{
    return 0, errors.New("ошибка")
  }
  return a/b, errors.New("ошибка ")
}