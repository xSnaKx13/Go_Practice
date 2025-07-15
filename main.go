package main

import (
	"fmt"
	"unicode"
)

func main(){
  str := "Правило никто не отменял!"
  
  for _, elem := range str{
    if unicode.IsUpper(elem[0]){
      fmt.Println("да")
    }else{
      fmt.Println("нет")
    }
  }
}