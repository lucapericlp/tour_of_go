package main

import (
	"fmt"
	"math/cmplx"
)

//variable declarations can be factored into blocks
var (
	ToBe	bool		= false
	maxInt	uint64		= 1 << 64 -1
	z		complex128	= cmplx.Sqrt(-5 + 12i)
)

func main(){
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
