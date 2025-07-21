package main

import (
	"fmt"
)

func main(){
  devide(10,2)
}

func devide(a,b int) (int, error){
  fmt.Println("Введите два целых числа:")
  _,error := fmt.Scan(&a, &b)
  if error != nil{
    error = fmt.Println("ошибка")
  }else{
    devide(a, b)
  }
  return error, a/b
}