package main

import (
	"fmt"
	"strings"
)

func main() {

	IP := "10.192.168.10"
	IP_L := strings.Split(IP, ".")

	if len(IP_L) == 4 {
		fmt.Println("IP has correct oct value")
		var x int = 1
		for _, value := range IP_L {
			if x == 1 {
				fmt.Printf("%v is the first oct of your IP \n", value)
			} else if x == 2 {
				fmt.Printf("%v is the second oct of your IP \n", value)
			} else if x == 3 {
				fmt.Printf("%v is the thrid oct of your IP \n", value)
			} else if x == 4 {
				fmt.Printf("%v is the fourth oct of your IP \n", value)
			}
			x++
		}

	}

}
