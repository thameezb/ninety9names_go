package models

//Names acts as an in memory cache of the list of names
var Names []Name

//Name represents one of the 99 names
type Name struct {
	Arabic          string
	Transliteration string
	MeaningShaykh   string
	Meaning         string
	Explanation     string
}
