package main

import (
	"fmt"
)

func main() {

	var username, password string

	fmt.Println("Please enter the username :")
	fmt.Scanf("%v\n%v", &username, &password)
	fmt.Println("Given username and password is : ", username, "and", password)
}
