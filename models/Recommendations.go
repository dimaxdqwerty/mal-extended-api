package models

type Recommendations struct {
	Anime               Anime `json:"node"`
	NumRecommendations int  `json:"num_recommendations"`
}
