package main

import "fmt"

func main() {
   var a,b,c,d,f string
   fmt.Scan(&a, &b, &c, &d, &f)
   names := []string {a,b,c,d,f}
      slice1 := names[:3]
	     fmt.Print(slice1)
		 }
