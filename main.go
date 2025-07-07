package main

import (
	"fmt"
)

func main(){
   minimumFromFour()
}
func minimumFromFour() int{
   var minValue int
   var number int
   fmt.Scan(&number)
   minValue = number
   for i := 1; i < 4; i++ {
      fmt.Scan(&number)
      if number < minValue{
         minValue = number
      }
   }
   fmt.Print(minValue)
   return minValue
}