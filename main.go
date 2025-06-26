package main

import "fmt"

func main() {
   var variable int
   fmt.Scan(&variable)

   number1 := variable / 100
   number2 := (variable / 10) % 10
   number3 := variable % 10

   numbers := fmt.Sprintf("%d%d%d", number3, number2, number1)
   fmt.Print(numbers)
}
   
