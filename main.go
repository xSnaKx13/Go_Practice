package main

import (
  "fmt"
  "math"
)

func main(){
  var weight float64
  var height float64 
  fmt.Scan(&weight)
  fmt.Scan(&height)
  IMT := weight / math.Pow(height, 2)
  fmt.Println(IMT)
}
