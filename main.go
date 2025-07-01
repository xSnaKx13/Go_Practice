package main

import "fmt"

func main() {
   var firstNumber, secondNumber int
   fmt.Scan(&firstNumber, &secondNumber)

   var result int

   for i := firstNumber; i < secondNumber; i++ {
      if i % 7 == 0{
         result = i
      }else{
         fmt.Print("NO")
      }
   }
   fmt.Print(result)
}
   
