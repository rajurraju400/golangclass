package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("rr-new.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Welco")
	file.Close()
}
