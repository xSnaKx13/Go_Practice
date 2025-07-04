package main

import (
	"fmt"
)

func main(){
   var a, b int
   fmt.Scan(&a,&b)

   var nums []int
   var nums2 []int
   var dublicateNums []int

   for a != 0 {
      value := a % 10
      nums = append(nums, value)
      a /= 10
   }

   for b != 0{
      value := b%10
      nums2 = append(nums2, value)
      b /= 10
   }
   for i := len(nums)-1; i >= 0; i-- {
      fmt.Print(nums[i])
   }
   fmt.Print(" ")
   for i := len(nums2)-1; i >= 0; i-- {
      fmt.Print(nums2[i])
   }

   for i := 0; i < len(nums); i++ {
      for j := 0; j < len(nums2); j++ {
         for k := 0; k < len(dublicateNums); k++ {
            if nums[i] == nums2[i]{
               if dublicateNums[k]!= nums[i] {
                  dublicateNums = append(dublicateNums, nums[i])
               }
            }
         }
      }
   }
   fmt.Println()
   for i := len(dublicateNums)-1; i >= 0; i-- {
      fmt.Print(dublicateNums[i], " ")
   }
}