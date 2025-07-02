package main

import (
	"fmt"
)

func main() {

   var numb int
   fmt.Scan(&numb)
   a, b := 0,1
   c, count := 0,1

   for {
      c = b + a
      count++
      if c == numb{
         fmt.Println(count)
         break
      }
      a = c + b
      count++
      if a == numb{
         fmt.Println(count)
         break
      }
      b = a + c
      count++
      if b == numb{
         fmt.Println(count)
         break
      }
   }
}
   
