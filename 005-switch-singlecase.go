package main

import (
	"fmt"
)

func main() {
	var day int = 8

	switch day {
	case 1:
		fmt.Println("This is number 1")
	case 2:
		fmt.Println("This is number 2")
	case 3:
		fmt.Println("This is number 3")
	default:
		fmt.Println("This numnber is unknow")
	}
}
