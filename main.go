package main

import "fmt"

func main() {
   var variable int
   fmt.Scan(&variable)
   arr := make([]int, variable)

   for i := 0; i < variable; i++ {
      fmt.Scan(&arr[i])
   }

   for idx, elem := range arr{
      if idx % 2 == 0{
         fmt.Println(elem)
      }
   }
}
   
