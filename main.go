package main

import (
	"fmt"
)

func main(){
   fibonacci(5)
}
func fibonacci(n int) int{
   fib := []int{0,1}
   for i := 0; i <= n; i++ {
      if i>=2{
      fib = append(fib, (fib[i-1]+fib[i-2]))
      }
   }
   fmt.Print(fib[n])
   return fib[n]
}