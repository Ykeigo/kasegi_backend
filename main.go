package main

import (
	//"database/sql"
	//"fmt"

	"github.com/gin-gonic/gin"
	//"github.com/lib/pq"
)

func main() {
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
/*
func createUser(models.user) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "password", "example", //TODO: 本番環境の値を入れる
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	v.Insert()
	defer db.Close()
}
*/