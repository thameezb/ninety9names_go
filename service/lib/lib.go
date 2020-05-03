package lib

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//ConverttoCSV reads in an excel file from readPath, converts sheet to csv stored at writePath
func ConvertToCSV(readPath, writePath, sheet string) error {
	file, err := excelize.OpenFile("orig.xlsx")
	if err != nil {
		return err
	}
	log.Print("file opened")

	var data []string
	for _, row := range file.GetRows(sheet) {
		if strings.Join(row, "") == "" {
			continue
		}
		data = append(data, strings.Join(row[1:], ","))
	}
	log.Print("data read into memory")

	ioutil.WriteFile(writePath, []byte(strings.Join(data, "\n")), 0644)
	log.Printf("file created succesfully at %s", writePath)
	return nil
}
