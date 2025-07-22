package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
  var input string
  input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
  input = strings.TrimSpace(input)

  var numbers[] int
  for i := 0; i < len(input); i++ {
    s, _:= strconv.Atoi(string(input[i]))
    numbers = append(numbers, s)
  }
  var maxNumber int = 0
  for _, elem := range numbers {
    if elem > maxNumber{
      maxNumber = elem
    }
  }
  fmt.Println(maxNumber)
}
