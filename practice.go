package main
import "fmt"

func dude(){
	var countNumber int
	fmt.Scan(&countNumber)
	var num int
	var positiveNumber int

	for i := 0; i < countNumber; i++ {
		fmt.Scan(&num)
		if num > 0{
			positiveNumber++
		}
	}
	fmt.Print(positiveNumber)
}