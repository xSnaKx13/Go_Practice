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

  for i := len(input)-1; i > 0; i-- {
    fmt.Printf("%s", input[i])
  }
}