package main

import (
	"fmt"
	"time"
)

//switch true is useful to clean up long if else chains
func main(){
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("Gm")
	case t.Hour() < 17:
		fmt.Println("Ga")
	default:
		fmt.Println("Gn")
	}
}
