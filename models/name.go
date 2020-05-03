package models

//Names acts as an in memory cache of the list of names
var Names map[int]Name

//Name represents one of the 99 names
type Name struct {
	Arabic          string `json:"arabic"`
	Transliteration string `json:"transliteration"`
	MeaningShaykh   string `json:"meaning_shaykh"`
	Meaning         string `json:"meaning_general"`
	Explanation     string `json:"explanation"`
}
