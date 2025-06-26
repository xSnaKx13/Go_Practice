package main

import "fmt"

func main() {
   array := [5]int{}
   max := 0
   for i := 0; i < 5; i++ {
      fmt.Scan(&array[i])
      if array[i] > max{
         max = array[i]
      }
   }
   fmt.Print(max)
}