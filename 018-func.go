package main

import (
	"fmt"
)

func main() {

	x, y := 10, 20

	fmt.Println("Addition of given two variables :", add(x, y))
	fmt.Println("sub of given two variables :", add(x, y))
	fmt.Println("multiplication of given two variables :", add(x, y))
}

func add(num1 int, num2 int) (int int) {

	var num3 = num1 + num2

	return num3

}

func sub(num1 int, num2 int) (int int) {

	var num3 = num1 - num2

	return num3

}

func mul(num1 int, num2 int) (int int) {

	var num3 = num1 * num2

	return num3

}
