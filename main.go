package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main(){
  input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  input = strings.TrimSpace(input)

  inputRune := []rune(input)
  inputFirstRune := []rune(input)[0]
  inputLastRune := inputRune[len(inputRune)-1]

  if unicode.IsUpper(inputFirstRune) && inputLastRune == '.'{
    fmt.Println("Right")
  }else{
    fmt.Println("Wrong")
  }
}