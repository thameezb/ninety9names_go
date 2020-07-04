package models

//Name represents one of the 99 names
type Name struct {
	ID              string `db:"id"                 json:"id"`
	Arabic          string `db:"arabic"             json:"arabic"`
	Transliteration string `db:"transliteration"    json:"transliteration"`
	MeaningShaykh   string `db:"meaning_shaykh"     json:"meaning_shaykh"`
	MeaningGeneral  string `db:"meaning_general"    json:"meaning_general"`
	Explanation     string `db:"explanation"        json:"explanation"`
}
