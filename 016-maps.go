package main

import (
	"fmt"
)

func main() {

	StudentRecord := make(map[string]int)

	test := []int{}

	PP := append(test, 1, 2, 3)

	StudentRecord["raj"] = 10
	StudentRecord["suresh"] = 12

	fmt.Println(len(StudentRecord))
	fmt.Println(StudentRecord["raj"])
	fmt.Println(int(PP[0]))

}
