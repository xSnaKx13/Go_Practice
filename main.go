package main

import (
	"fmt"
)

func main(){
   var a int
   fmt.Scan(&a)

   var nums []int

   for {
      value := a % 10
      nums = append(nums, value)
      value /= 10
      if value == 0{
         break
      }
   }

   fmt.Print(nums, " ")
}