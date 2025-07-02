package main

import (
	"fmt"
)

func main() {
   a := 73081
   b := 650
   var result string

   numb := fmt.Sprintf("%d", a)
   numb2 := fmt.Sprintf("%d", b)

   fmt.Println(numb, " ", numb2)

   for i:=0; i<len(numb);i++{
      for k:=0; k<len(numb2);k++{
         if string(numb[i]) == string(numb2[k]){
            result += string(numb[i])
         }
      }
   }
   fmt.Printf(result)

}
   
