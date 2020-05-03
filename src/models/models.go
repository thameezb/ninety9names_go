package models

//Names acts as an in memory cache of the list of names
var Names map[int]Name

//Name represents one of the 99 names
type Name struct {
	ID              string `db:"id" 				json:"id"`
	Arabic          string `db:"arabic" 			json:"arabic"`
	Transliteration string `db:"transliteration" 	json:"transliteration"`
	MeaningShaykh   string `db:"meaningShaykh" 		json:"meaning_shaykh"`
	MeaningGeneral  string `db:"meaningGeneral" 	json:"meaning_general"`
	Explanation     string `db:"explanation" 		json:"explanation"`
}
