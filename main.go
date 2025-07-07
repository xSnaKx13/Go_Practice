package main

import (
   "fmt"
)

func main(){
   vote(5,0,5)
}
func vote(x,y,z int) int{
   result := 0
   switch {
   case x == y || x == z:
      result = x
   case y == x || y == z:
      result = y
   case z == x || z == y:
      result = z
   }
   fmt.Print(result)
   return result
}