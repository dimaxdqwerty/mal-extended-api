package operations

import (
	"encoding/json"
	"mal-extendet-api/models"
)

func GetGenres() []models.Genres {
	genresLen, err := client.ZCard("genres").Result()
	handleErr(err)
	result, err := client.ZRange("genres", 0, genresLen).Result()
	handleErr(err)
	var genres []models.Genres
	for _, str := range result {
		var genre models.Genres
		err = json.Unmarshal([]byte(str), &genre)
		handleErr(err)
		genres = append(genres, genre)
	}
	return genres
}
