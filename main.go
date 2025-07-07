package main

import (
   "fmt"
)

func main(){
   fibonacci(6)
}
func fibonacci(n int) int{
   var count int = 2
   var a int = 1
   var b int = 1
   var c int
   var value int

   for {
      if n <= count && n > 0{
         value = 1
         break
      }
      c = b + a
      count++
      if count == n{
         value = c
         break
      }
      a = c + b
      count++
      if count == n{
         value = a
         break
      }
      b = a + c
      count++
      if count == n{
         value = b
         break
      }
   }
   fmt.Print(value)
   return value
}