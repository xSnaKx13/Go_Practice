package main

import "fmt"

func main() {
   var variable int
   fmt.Scan(&variable)

   number1 := variable / 100
   number2 := (variable / 10) % 10
   number3 := variable % 10

   sumNumbers := number1 + number2 + number3
   fmt.Print(sumNumbers, " ")
}
   
