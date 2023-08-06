package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	models "test-app/models_psql"

	_ "github.com/lib/pq" //なんかよくわからんけどいる　ないとruntimeErrorを吐く
)

func insertTest() {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568
	db, _ := sql.Open(
		"postgres",
		"host=localhost port=5432 dbname=example user=postgres password=password sslmode=disable")
	titles, _ := models.GameTitles().All(context.Background(), db)
	fmt.Println("jets:", titles)
	for _, v := range titles {
		fmt.Println(v)
	}
}

func webServerTest() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})

	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	r.Run(":80")
}

func main() {
	webServerTest()
}
