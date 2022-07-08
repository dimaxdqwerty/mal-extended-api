package models

type AlternativeTitles struct {
	Synonyms []string `json:"synonyms"`
	En       string   `json:"en"`
	Ja       string   `json:"ja"`
}
