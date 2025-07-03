package main

import (
	"fmt"
)

func main() {
   var firstNumber int
   var secondNumber int

   fmt.Scan(&firstNumber, &secondNumber)

   var result string
   isDublicate := false

   numb := fmt.Sprintf("%d", firstNumber)
   numb2 := fmt.Sprintf("%d", secondNumber)

   fmt.Println(numb, " ", numb2)

   for i:=0; i<len(numb);i++{
      for k:=0; k<len(numb2);k++{
         if numb[i] == numb2[k]{
            for j := 0; j<len(result); j++{
               if numb[i] == result[j]{
                  isDublicate = true
                  break
               }
            }
            if !isDublicate{
               result += string(numb[i])
               result += " "
            }
         }
      }
   }
   fmt.Print(result)
}