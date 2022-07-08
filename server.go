package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"mal-extendet-api/db"
	"mal-extendet-api/models"
	"mal-extendet-api/operations"
	"net/http"
)

var router *gin.Engine
var routerGroup *gin.RouterGroup
var client = db.GetRedisClient()

func main() {
	router = gin.Default()

	result, err := client.Keys("*").Result()
	handleErr(err)
	
	if len(result) == 0 {
		dumpAnimeListJob()
	}

	initializeRoutes()

	go func() {
		err := gocron.Every(2).Hours().Do(dumpAnimeListJob)
		handleErr(err)
		<-gocron.Start()
	}()

	err = router.Run()
	handleErr(err)
}

func initializeRoutes() {
	routerGroup = router.Group("/api")

	routerGroup.GET("/anime/:id", func(context *gin.Context) {
		id := context.Param("id")
		anime := operations.GetAnimeByID(id)
		context.JSON(http.StatusOK, anime)
	})
	routerGroup.GET("/anime/ranking", func(context *gin.Context) {
		rankingType := context.Query("ranking_type")
		limit := context.Query("limit")
		offset := context.Query("offset")
		if len(limit) == 0 || len(offset) == 0 {
			context.JSON(http.StatusBadRequest, "Invalid Parameters")
			return
		}
		if len(rankingType) == 0 {
			rankingType = "rating"
		}

		var list []models.Anime
		var paging string

		switch rankingType {
		case "rating_desc":
			list, paging = operations.GetAnimeListWithParameters("animeListByRatingDesc", limit, offset)
		case "rating_asc":
			list, paging = operations.GetAnimeListWithParameters("animeListByRatingAsc", limit, offset)
		case "popularity_desc":
			list, paging = operations.GetAnimeListWithParameters("animeListByPopularityDesc", limit, offset)
		case "popularity_asc":
			list, paging = operations.GetAnimeListWithParameters("animeListByPopularityAsc", limit, offset)
		}
		context.JSON(
			http.StatusOK,
			gin.H{
				"paging":  paging,
				"ranking": list,
			})
	})
	
	routerGroup.GET("/genres", func(context *gin.Context) {
		genres := operations.GetGenres()
		context.JSON(http.StatusOK, genres)
	})
	routerGroup.GET("/studios", func(context *gin.Context) {
		studios := operations.GetStudios()
		context.JSON(http.StatusOK, studios)
	})
}

func dumpAnimeListJob() {
	fmt.Println("Starting dumping...")
	client.FlushAll()
	fmt.Println("Flushed all data!")
	fmt.Println("Staring dumping...")
	operations.DumpAnimeList()
	fmt.Println("Successfully dumped!")
	_, time := gocron.NextRun()
	fmt.Println("Next dump in ", time)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
