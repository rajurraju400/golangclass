package main

import (
	"fmt"
)

func main() {

	day := 10

	switch day {

	case 1, 2, 3:
		fmt.Println("Given number is 3")
	case 4, 5:
		fmt.Println("Given number is 5")
	}
}
