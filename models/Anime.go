package models

import (
	"time"
)

type Anime struct {
	ID                     int               `json:"id"`
	Title                  string            `json:"title"`
	MainPicture            MainPicture       `json:"main_picture"`
	AlternativeTitles      AlternativeTitles `json:"alternative_titles"`
	StartDate              string            `json:"start_date"`
	EndDate                string            `json:"end_date"`
	Synopsis               string            `json:"synopsis"`
	Mean                   float64           `json:"mean"`
	Rank                   int               `json:"rank"`
	Popularity             int               `json:"popularity"`
	NumListUsers           int               `json:"num_list_users"`
	NumScoringUsers        int               `json:"num_scoring_users"`
	Nsfw                   string            `json:"nsfw"`
	CreatedAt              time.Time         `json:"created_at"`
	UpdatedAt              time.Time         `json:"updated_at"`
	MediaType              string            `json:"media_type"`
	Status                 string            `json:"status"`
	Genres                 []Genres          `json:"genres"`
	MyListStatus           MyListStatus      `json:"my_list_status"`
	NumEpisodes            int               `json:"num_episodes"`
	StartSeason            StartSeason       `json:"start_season"`
	Broadcast              Broadcast         `json:"broadcast"`
	Source                 string            `json:"source"`
	AverageEpisodeDuration int               `json:"average_episode_duration"`
	Rating                 string            `json:"rating"`
	Pictures               []Pictures        `json:"pictures"`
	Background             string            `json:"background"`
	RelatedAnime           []RelatedAnime    `json:"related_anime"`
	//RelatedManga    		[]RelatedManga 		`json:"related_manga"` TODO: create RelatedManga struct
	Recommendations []Recommendations `json:"recommendations"`
	Studios         []Studios         `json:"studios"`
	Statistics      Statistics        `json:"statistics"`
}
