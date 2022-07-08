package models

type RelatedAnime struct {
	Anime                 Anime  `json:"node"`
	RelationType          string `json:"relation_type"`
	RelationTypeFormatted string `json:"relation_type_formatted"`
}
