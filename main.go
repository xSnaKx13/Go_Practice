package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
  var input string
  input, _ = bufio.NewReader(os.Stdin).ReadString('\n')

  number, _ := strconv.Atoi(input)
  var sliceNumbs[] int

  for i := len(input)-1; i > 0; i-- {
    s := number % 10
    sliceNumbs = append(sliceNumbs, s)
  }
  fmt.Println(sliceNumbs)
}
