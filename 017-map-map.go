package main

import (
	"fmt"
)

func main() {

	superhero := map[string]map[string]string{

		"batman": {
			"realname": "raj",
			"city":     "gotham city",
		},
		"suresh": {
			"realname": "raj",
			"city":     "gotham city",
		},
	}
	fmt.Println(superhero["batman"])
}
