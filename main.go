package main

import (
	"fmt"
)

func main() {
   var firstNumb, secondNumb int
   fmt.Scan(&firstNumb,&secondNumb)

   slice := make([]int, firstNumb)
   slice2 := make([]int, secondNumb)

   for i := 0; i < len(slice); i++ {
      fmt.Println(slice[i])
   }
   for i := 0; i < len(slice2); i++ {
      fmt.Println(slice2[i])
   }
}
   
