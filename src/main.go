// https://github.com/zalando/gin-oauth2/blob/master/google/google.go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	models "test-app/models_psql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/v4/boil"

	_ "github.com/lib/pq" //なんかよくわからんけどいる　これがないとDBアクセスがruntimeErrorを吐く
)

func insertTest(user models.User, db *sql.DB) {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568
	err := user.Insert(context.Background(), db, boil.Infer())
	fmt.Println(err)
}

func selectTest(db *sql.DB) {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	user, e := models.FindUser(context.Background(), db, "")
	fmt.Println(user)
	fmt.Println(e)
}

func webServerTest(google *Google) {
	
	db, e := sql.Open(
		"postgres",
		"host=localhost port=5432 dbname=example user=postgres password=password sslmode=disable")
	if e != nil {
		fmt.Println(e)
		return
	}
	defer db.Close()

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
	r.GET("/auth", func(c *gin.Context) {
		var redirectUrl = google.GetLoginURL("state")// todo: ランダムなstateを生成する
		fmt.Println(redirectUrl)
		c.JSON(200, gin.H{
			"redirect-url": redirectUrl,
		})
	})
	r.GET("/google/callback", func(c *gin.Context) {
		var userId, error = google.GetUserID(c.Query("code"))
		if error != nil {
			c.JSON(500, gin.H{
				"message": error.Error(),
			})
		} else {
			//todo:アカウントがなければつくる

			//todo:セッションを発行する
			c.JSON(200, gin.H{
				"message": "login completed",
				"userId" : userId,
			})
		}
	})
	r.GET("/insert", func(c *gin.Context) {
		insertTest(models.User{UserName: "test"}, db)
		c.JSON(200, gin.H{
		"message": "hello",
		})
	})
	r.GET("/select", func(c *gin.Context) {
		selectTest(db)
		c.JSON(200, gin.H{
		"message": "hello",
		})
	})
	r.Run(":8080")
}

func loadEnv() {
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load("../oidc_credentials.env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// .envの SAMPLE_MESSAGEを取得して、messageに代入します。
	message := os.Getenv("GoogleClientID")

	fmt.Println(message)
}

func main() {
	loadEnv()
	var google = NewGoogle()
	webServerTest(google)
}
