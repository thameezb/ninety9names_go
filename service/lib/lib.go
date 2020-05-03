package lib

import (
	"io/ioutil"
	"log"
	"ninety9names/models"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//ConverttoCSV reads in an excel file from readPath, converts sheet to csv stored at writePath
func ConvertToCSV(readPath, writePath, sheet string) error {
	file, err := excelize.OpenFile("orig.xlsx")
	if err != nil {
		return err
	}

	var data []string
	for _, row := range file.GetRows(sheet) {
		if strings.Join(row, "") == "" {
			continue
		}
		data = append(data, strings.Join(row[1:], ","))
	}

	ioutil.WriteFile(writePath, []byte(strings.Join(data, "\n")), 0644)
	log.Printf("file created successfully at %s", writePath)
	return nil
}

//SetNames reads input from csv and creates names map
func SetNames(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	for i, l := range strings.Split(string(data), "\n")[1:] {
		e := strings.Split(l, ",")
		models.Names[i] = models.Name{
			Arabic:          e[0],
			Transliteration: e[1],
			MeaningShaykh:   e[2],
			Meaning:         e[3],
			Explanation:     e[4],
		}
	}
	return nil
}
