package main

import (
	"fmt"
)

func main(){
  str := "правило никто не отменял!"

  lastS := []rune(str)[len(str)]
  if lastS == '!'{
    fmt.Println("да")
  }else{
    fmt.Println("нет")
  }
}