package main

import (
	"fmt"
)

func main(){
   var a int
   fmt.Scan(&a)

   var nums []int

   for a != 0 {
      value := a % 10
      nums = append(nums, value)
      a /= 10
   }

   fmt.Println(nums, " ")

   for i := len(nums); i > len(nums); i-- {
      fmt.Println(nums[i])
   }
}