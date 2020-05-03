package main

import (
	"log"
	"ninety9names/service/lib"
)

func main() {
	err := lib.ConvertToCSV("orig.xlsx", "./input/input.csv", "Sheet1")
	if err != nil {
		log.Print(err)
	}
}
