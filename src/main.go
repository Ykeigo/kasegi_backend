// https://github.com/zalando/gin-oauth2/blob/master/google/google.go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	api "kasegi/api"
	models "kasegi/my_models"
	repository "kasegi/repository"
	"kasegi/util"

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

	userRepository := repository.UserRepository{}
	loginSessionRepository := repository.LoginSessionRepository{}
	gameMatchApi := api.GameMatchApi{}
	
	db, e := sql.Open(
		"postgres",
		"host=localhost port=5432 dbname=kasegi user=postgres password=password sslmode=disable")
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
		var email, error = google.VerifyAndGetEmail(c.Query("code"))
		if error != nil {
			c.JSON(500, gin.H{
				"message": error.Error(),
			})
		} else {
			//アカウントがなければつくる
			var existUser = userRepository.FindByEmail(email, db)
			if existUser == nil {
				userRepository.Insert(models.User{Email: email}, db)
				existUser = userRepository.FindByEmail(email, db)
			}
			//todo:セッションを発行する
			var SessionToken = util.IdGenerator{}.GenerateSurrogateKey()
			loginSessionRepository.DeleteByUserId(existUser[0].ID, db)
			loginSessionRepository.Insert(models.LoginSession{UserID: existUser[0].ID, SessionToken: SessionToken}, db)
			
			c.JSON(200, gin.H{
				"message": "login completed",
				"email" : email,
				"sessionToken" : SessionToken,
			})
		}
	})
	r.GET("/checkSession", func(c *gin.Context) {
		var sessionToken = c.Query("sessionToken")
		var loginSession = loginSessionRepository.FindBySessionToken(sessionToken, db)
		fmt.Println(loginSession)
		// 期限切れじゃないやつだけにする
		var validLoginSession []models.LoginSession
		for _, l := range loginSession {
			if l.ExpiresAt.After(time.Now()) {
				validLoginSession = append(validLoginSession, l)
			}
		}
		if len(validLoginSession) > 0 {
			c.JSON(200, gin.H{
				"message": "session is valid",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "session is not valid",
			})
		}
	})/*
	r.GET("/insert", func(c *gin.Context) {
		userRepository.Insert(models.User{GoogleUserID: c.Query("googleUserId")}, db)
		c.JSON(200, gin.H{
		"message": "hello",
		})
	})*/
	r.GET("/listUser", func(c *gin.Context) {
		var users = userRepository.List(db)
		c.JSON(200, gin.H{
		"message": users,
		})
	})
	r.POST("/createMyGameMatch", func(c *gin.Context) {
		gameMatchApi.CreateMyGameMatch(c, db)
	})	
	r.POST("/listMyGameMatch", func(c *gin.Context) {
		gameMatchApi.ListMyGameMatch(c, db)
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
