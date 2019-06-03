package main

import "fmt"

//we use this notation to document the meaning of the return values 
func split(sum int) (x,y int){
	x = sum * 4 / 9
	y = sum - x
	//known as a naked return
	//should be used only in short functions (readability)
	return
}

func main(){
	fmt.Println(split(17))
}
