package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
  var input string
  input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
  input = strings.TrimSpace(input)
  numbers := []rune(input)
  
  var numsRes[] int
  for _, elem := range numbers{
    digit := (elem - '0')
    sqr := digit * digit
    numsRes = append(numsRes, int(sqr))
  }
  for _, elem := range numsRes{
    fmt.Print(elem)
  }
}
