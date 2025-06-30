package main

import "fmt"

func main() {
   var nCount int
   fmt.Scan(&nCount)
   zeroCount := 0

   var numbers = make([]int, nCount)

   for i := 0; i < nCount; i++ {
      fmt.Scan(&numbers[i])
      if numbers[i] == 0 {
         zeroCount++
      }
   }
   fmt.Print(zeroCount)
}
   
