package main

import (
	"fmt"
	"strings"
)

func main() {
	var num string = "        Welcome to chennai"

	num1 := strings.TrimSpace(num)
	fmt.Println(num)
	fmt.Println(num1)
}
