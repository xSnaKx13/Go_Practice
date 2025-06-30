package main

import "fmt"

func main() {
   var number int
   fmt.Scan(&number)
   var testNum string

   for number > 0{
      digit := number%10
      if digit != 7{
         testNum += fmt.Sprintf("%d", digit)
      }
      number /= 10
   }
   fmt.Print(testNum)
   
}
   
