package main

import (
	"fmt"
)

func main(){
   var a int
   fmt.Scan(&a)

   var nums []int

   nums = append(nums, a)

   fmt.Print(nums)
}