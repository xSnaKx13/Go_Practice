package main

import "fmt"

func main() {
   var number int
   fmt.Scan(&number)
   sumNumbers := 0

   for number > 0{
      digit := number % 10
      sumNumbers += digit
      number /= 10
      
   }
   fmt.Print(sumNumbers)
}
   
