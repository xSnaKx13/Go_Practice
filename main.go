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
    sqr := elem * elem
    numsRes = append(numsRes, int(sqr))
  }
  fmt.Print(numsRes)
}
