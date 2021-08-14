package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	ii := "123"

	i, err := strconv.Atoi(ii)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
	fmt.Printf("datatype of the i is %T \n", i)

}
