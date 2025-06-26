package main

import "fmt"

func main() {
   var variable int
   fmt.Scan(&variable)
   
   slice := make([]int, variable)
   var value int
   for i := 0; i < variable; i++{
      fmt.Scan(&value)
      slice[i] = value
   }
   fmt.Print(slice[3])
}