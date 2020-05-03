package main

import (
	"log"
	"ninety9names/service/lib"
)

func readData() {
	err := lib.ConvertToCSV("orig.xlsx", "./input/input.csv", "Names")
	if err != nil {
		log.Print(err)
	}

	err = lib.SetNames("./input/input.csv")
	if err != nil {
		log.Print(err)
	}
}

func main() {
	readData()
}
