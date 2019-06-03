package main

import "fmt"

func main(){
	fmt.Println("counting")

	//defer works by pushing function calls onto a stack
	for i := 0; i < 10; i++{
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
