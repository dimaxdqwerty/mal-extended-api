package operations

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"mal-extendet-api/db"
	"mal-extendet-api/models"
	"net/http"
	"sort"
	"strconv"
)

type AnimeList struct {
	Data   []models.Data `json:"data"`
	Paging models.Paging `json:"paging"`
}

var (
	MalClientID = models.GetMalClientID()

	GetAnimeQuery = "https://api.myanimelist.net/v2/anime"

	Fields = "fields=id,title,main_picture,alternative_titles,start_date,end_date," +
		"synopsis,mean,rank,popularity,num_list_users,num_scoring_users,nsfw,created_at," +
		"updated_at,media_type,status,genres,my_list_status,num_episodes,start_season," +
		"broadcast,source,average_episode_duration,rating,pictures,background,related_anime," +
		"recommendations,studios,statistics" //TODO: add related_manga field when RelatedManga struct will be added
)

var client = db.GetRedisClient()

func GetAnimeByID(ID string) models.Anime {
	animeBinary, err := client.Get(ID).Result()
	handleErr(err)
	var anime models.Anime

	err = json.Unmarshal([]byte(animeBinary), &anime)
	handleErr(err)
	return anime
}

func GetAnimeRankingList(limit string, offset string) AnimeList {
	var animeList AnimeList
	req, err := http.NewRequest("GET", GetAnimeQuery+"/ranking"+"?rankingType=all"+"&limit="+limit+"&offset="+offset+"&"+Fields, nil)
	req.Header.Add("X-MAL-CLIENT-ID", MalClientID)
	handleErr(err)

	client := &http.Client{}
	resp, err := client.Do(req)
	handleErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	handleErr(err)

	err = json.Unmarshal(body, &animeList)
	handleErr(err)
	return animeList
}

func GetAnimeRankingListViaPaging(paging models.Paging) AnimeList {
	var listViaPaging AnimeList
	req, err := http.NewRequest("GET", paging.Next, nil)
	req.Header.Add("X-MAL-CLIENT-ID", MalClientID)
	handleErr(err)

	client := &http.Client{}
	resp, err := client.Do(req)
	handleErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	handleErr(err)

	err = json.Unmarshal(body, &listViaPaging)
	handleErr(err)
	return listViaPaging
}

func GetWholeAnimeList() []AnimeList {
	var list []AnimeList
	list = append(list, GetAnimeRankingList("500", "0"))
	for list[len(list)-1].Paging.Next != "" {
		list = append(list, GetAnimeRankingListViaPaging(list[len(list)-1].Paging))
	}
	return list
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetAnimeListWithParameters(rankingType string, limit string, offset string) ([]models.Anime, string) {
	var animeList []models.Anime
	
	lim, err := strconv.Atoi(limit)
	handleErr(err)
	
	off, err := strconv.Atoi(offset)
	handleErr(err)
	
	result, err := client.LRange(rankingType, int64(off), int64(off)+int64(lim)).Result()
	handleErr(err)
	
	for _, animeBinary := range result {
		var anime models.Anime
		err = json.Unmarshal([]byte(animeBinary), &anime)
		handleErr(err)
			animeList = append(animeList, anime)
	}
	
	paging := "/api/ranking/limit?=" + strconv.Itoa(lim) + "&offset=" + strconv.Itoa(off+lim)
	return animeList, paging
}

func GetAnimeListLen() int64 {
	animeListLen, err := client.LLen("animeListByRatingDesc").Result()
	handleErr(err)
	return animeListLen
}

func MarshalBinary(anime interface{}) ([]byte, error) {
	return json.Marshal(anime)
}

func DumpAnimeList() {
	list := GetWholeAnimeList()
	var dataList []models.Anime
	for _, dataArray := range list {
		for _, data := range dataArray.Data {
			anime, err := MarshalBinary(data.Anime)
			handleErr(err)
			dataList = append(dataList, data.Anime)
			client.Set(strconv.Itoa(data.Anime.ID), anime, 0)
			client.RPush("animeListByRatingDesc", anime)
			client.LPush("animeListByRatingAsc", anime)
			for _, genre := range data.Anime.Genres {
				genreBinary, err := MarshalBinary(genre)
				handleErr(err)
				client.ZAdd("genres", redis.Z{Score: float64(genre.ID), Member: genreBinary})
			}
			for _, studio := range data.Anime.Studios {
				studioBinary, err := MarshalBinary(studio)
				handleErr(err)
				client.ZAdd("studios", redis.Z{Score: float64(studio.ID), Member: studioBinary})
			}
		}
	}
	sort.SliceStable(dataList, func(i, j int) bool {
		return dataList[i].Popularity < dataList[j].Popularity
	})
	for _, anime := range dataList {
		animeBinary, err := MarshalBinary(anime)
		handleErr(err)
		client.RPush("animeListByPopularityDesc", animeBinary)
	}
	sort.SliceStable(dataList, func(i, j int) bool {
		return dataList[i].Popularity > dataList[j].Popularity
	})
	for _, anime := range dataList {
		animeBinary, err := MarshalBinary(anime)
		handleErr(err)
		client.RPush("animeListByPopularityAsc", animeBinary)
	}
}
