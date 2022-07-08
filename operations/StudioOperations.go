package operations

import (
	"encoding/json"
	"mal-extendet-api/models"
)

func GetStudios() []models.Studios {
	studiosLen, err := client.ZCard("studios").Result()
	handleErr(err)
	result, err := client.ZRange("studios", 0, studiosLen).Result()
	handleErr(err)
	var studios []models.Studios
	for _, str := range result {
		var studio models.Studios
		err = json.Unmarshal([]byte(str), &studio)
		handleErr(err)
		studios = append(studios, studio)
	}
	return studios
}
