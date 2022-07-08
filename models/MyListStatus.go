package models

import "time"

type MyListStatus struct {
	Status             string    `json:"status"`
	Score              int       `json:"score"`
	NumEpisodesWatched int       `json:"num_episodes_watched"`
	IsRewatching       bool      `json:"is_rewatching"`
	UpdatedAt          time.Time `json:"updated_at"`
}
