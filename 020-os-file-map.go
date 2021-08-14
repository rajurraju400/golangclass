package main

import (
	"fmt"
)

func main() {
	StudentDetails := make(map[string]int)
	StudentDetails["raj"] = 10
	fmt.Println(StudentDetails["raj"])

	superhero := map[string]map[string]string{
		"batman": {
			"play":   "triger",
			"game":   "criket",
			"food":   "briyani",
			"cloths": "welcome",
		},
		"sure": {
			"play": "peek",
		},
	}

	fmt.Println(superhero["batman"]["game"])
}
