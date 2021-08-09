package main

import (
	"fmt"
)

func main() {
	var animal string = "Tiger"

	if animal == "Tiger" {
		fmt.Println("Name of the wild animal is :", animal)
	} else if animal == "cow" {
		fmt.Println("name of the corni animal is ", animal)
	} else {
		fmt.Println("name of the animal is unknow")
	}
}
