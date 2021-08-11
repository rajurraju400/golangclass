package main

import (
	"fmt"
)

func main() {
	var ip string = "192.168.0.10"

	fmt.Println("given IP is:", ip)

	num := []int{1, 2, 3, 4, 5}

	fmt.Printf("Value of num is : %v", num)

	num = append(num, 10, 20)

	fmt.Printf("\nDisplaying the updated num variable %v", num)
}
