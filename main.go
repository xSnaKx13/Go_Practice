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

  number, _ := strconv.Atoi(input)
  var sliceNums []int 

  for number >= 0{
    s := number % 10
    sliceNums = append(sliceNums, s)
  }
  fmt.Print(sliceNums)
}
