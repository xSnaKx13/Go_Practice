package main

import (
	"fmt"
	"strings"
)

func main(){
  var s string = "qwerty"
  newS := strings.Split(s, "")
  result := strings.Join(newS, "*")
  fmt.Println(result)
}
