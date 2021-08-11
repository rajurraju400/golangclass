package main

import (
	"fmt"
	"strings"
)

func main() {

	var test string = "Welcome to Chennai"

	fmt.Println(strings.Split(test, " ")[:2])

	tt := strings.Split(test, " ")
	fmt.Println(tt)
}
