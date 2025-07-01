package main

import "fmt"

func main() {
   var firstNumber, secondNumber int
   fmt.Scan(&firstNumber, &secondNumber)

   var result string

   for i := secondNumber; i >= firstNumber; i-- {
      if i % 7 == 0{
         result = fmt.Sprintf("%d", i)
         break
      }else{
         result = "NO"
      }
   }
   fmt.Print(result)
}
   
