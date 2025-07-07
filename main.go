package main

import (
	"fmt"
)

func main(){
   test()
}
func test(){
   foo := []interface{}{}
   for i := 0; foo[i] !=34; i++ {
      fmt.Scan(&foo[i])
   }
   fmt.Print(foo...)
}
