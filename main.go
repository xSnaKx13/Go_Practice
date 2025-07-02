package main

import (
	"fmt"
)

func main() {
   var firstNumb, secondNumb int
   fmt.Scan(&firstNumb,&secondNumb)
   var a int

   for firstNumb > 0 {
      a = firstNumb % 10
   }
   fmt.Print(a)
}
   
