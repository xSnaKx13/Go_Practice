package main

import (
	"fmt"
)

func main() {
   a := 730570
   b := 65095
   var result string

   numb := fmt.Sprintf("%d", a)
   numb2 := fmt.Sprintf("%d", b)

   fmt.Println(numb, " ", numb2)

   for i:=0; i<len(numb);i++{
      for k:=0; k<len(numb2);k++{
         if numb[i] == numb2[k]{
            for j := 0; j<len(result); i++{
               if numb[i] == result[j]{
                  continue
               }
            }
            result += string(numb[i])
         }
      }
   }
   fmt.Printf(result)

}
   
