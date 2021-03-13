package main

import (
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/thameezb/ninety9names/src/models"
	"github.com/thameezb/ninety9names/src/repository"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("$DATABASE_URL must be set")
	}
	db := repository.New(dbURL)

	filePath := os.Getenv("SOURCE_FILE")
	if filePath == "" {
		filePath = "orig.xlsx"
	}
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = db.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range file.GetRows("Names") {
		if row[0] == "#" {
			continue
		}
		err := db.InsertName(models.Name{
			ID:              row[0],
			Arabic:          row[1],
			Transliteration: row[2],
			MeaningShaykh:   row[3],
			MeaningGeneral:  row[4],
			Explanation:     row[5],
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Print("data inject complete")
}
