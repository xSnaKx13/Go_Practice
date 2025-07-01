package main

import "fmt"

func main() {
   var numb int
   fmt.Scan(&numb)

   cowCount := numb % 10

   switch cowCount{
   case 1:
      fmt.Printf("%d korova", numb)
   case 2, 3, 4:
      fmt.Printf("%d korovy", numb)
   case 0, 5, 6, 7, 8, 9:
      fmt.Printf("%d korov", numb)
   }

}
   
