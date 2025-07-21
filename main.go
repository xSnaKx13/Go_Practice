package main

import (
	"fmt"
	"math"
)

func main(){
  var a, b int
  _, err := fmt.Scan(&a, &b)
  if err != nil{
    fmt.Println("ошибка")
    return
  }
  hypotenuse(a, b)
}
func hypotenuse(a,b int)int{
    squareA := a * a
    squareB := b * b
    sumAB := squareA + squareB
    hptnse := math.Sqrt(float64(sumAB))
    fmt.Println(int(hptnse))
    return int(hptnse)
}
