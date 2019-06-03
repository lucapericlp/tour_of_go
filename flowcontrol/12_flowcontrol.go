package main

import "fmt"

func main(){
	defer fmt.Println("world") //executed when surrounding function returns
	//deferrred call's args are evalued immediately just not executed.	
	fmt.Println("hello")
}
