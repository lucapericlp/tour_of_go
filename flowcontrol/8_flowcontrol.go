package main

import (
	"fmt"
)

func Sqrt(x float64) float64{
	z := 1.0
	for i := 1; i < 11; i++{
		var delta float64 = (z*z-x) / (2*z)
		z -= delta
	}
	return z
}

func main(){
	fmt.Println(Sqrt(2))
}
