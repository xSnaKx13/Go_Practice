package main

import "fmt"

func main() {
   var numbers []int
   var nCount int
   fmt.Scan(&nCount)
   zeroCount := 0
   for i := 0; i < nCount; i++ {
      fmt.Scan(&numbers[i])
      if numbers[i] == 0 {
         zeroCount++
      }
   }
   fmt.Print(zeroCount)
}
   
