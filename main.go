package main

import (
	"fmt"
)

func main(){
  

}

func devide(a,b int) (resilt int){
  fmt.Println("Введите два целых числа:")
  _,err := fmt.Scan(&a, &b)
  if err != nil{
    fmt.Println("Введено некоректное знаяение")
  }else{
    result = a / b
  }
  return result
}