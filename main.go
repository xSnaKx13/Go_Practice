package main

import (
	"fmt"
)

func main(){
   sumInt(10,2,4,1,3,1,80,1)
}
func sumInt(numbers...int){
   var result int
   var countParametrs int
   for _, elem := range numbers{
      countParametrs++
      result += elem
   }
   fmt.Print(countParametrs, " ", result)
}
