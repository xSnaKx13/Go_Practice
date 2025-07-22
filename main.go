package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main(){
  var input string
  input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
  input = strings.TrimSpace(input)
  numbers := []rune(input)
  fmt.Println(string(slices.Max(numbers)))
}
