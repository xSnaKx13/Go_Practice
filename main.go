package main

import "fmt"

func main() {
   n := 273
   k := n % 10
   var res int
   switch res{
   case k == 1:
      fmt.Print("корова")
   case k == 2:
      fmt.Print("коровы")
   }
}
   
