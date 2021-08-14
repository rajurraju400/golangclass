package main

import (
	"fmt"
)

func main() {

	test01 := []string{"goat", "tiger", "cow", "Dear", "2"}

	for _, test02 := range test01 {
		if test02 == "goat" {
			fmt.Printf("%v is the given animal\n", test02)
		} else if test02 == "tiger" {
			test03 := append(test01, "dog")
			fmt.Println(len(test03))

		} else {
			fmt.Printf("not a goat on this time!\n")
		}

	}
}
