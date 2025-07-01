package main

import "fmt"

func main() {
   var numb int
   fmt.Scan(&numb)

   square := 1

   for square < numb{
      fmt.Print(square, " ")
      square *= 2
   }
}
   
