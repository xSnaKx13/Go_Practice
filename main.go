package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
  input,_ := bufio.NewReader(os.Stdin).ReadString('\n')
  input = strings.TrimSpace(input)

  word := []rune(input)
  var drow string

  for i := len(word) - 1; i >= 0; i-- {
    drow += string(word[i])
  }
  if input == drow{
    fmt.Print("Палиндром")
  }else{
    fmt.Print("Нет")
  }
}