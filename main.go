package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
  var input string
  input, err := bufio.NewReader(os.Stdin).ReadString(('\n'))
  input = strings.TrimSpace(input)
  if err != nil{
    fmt.Println("ошибка")
    return
  }
  newString := strings.Split(input, "")
  result := strings.Join(newString, "*")
  fmt.Println(result)
}
