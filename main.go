package main

import (
	"fmt"
)

type Cars struct{
  _brand string
  _model string
  _horsePower float64
  _yearRealize int
}

func NewCar(
  brand string,
  model string,
  horsePower float64,
  yearRealize int,
) Cars {
  if brand == "" {
    return Cars{}
  }
  if model == "" {
    return Cars{}
  }
  if horsePower < 70 || horsePower > 560 {
    return Cars{}
  }
  if yearRealize < 1990 || yearRealize > 2025{
    return Cars{}
  }
  return Cars{
    _brand: brand,
    _model: model,
    _horsePower: horsePower,
    _yearRealize: yearRealize,
  }
}

func main(){
  audi := NewCar("audi", "a4", 129.0, 2021)
  fmt.Println(audi)
}