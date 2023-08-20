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
	"golang.org/x/oauth2"

	_ "github.com/lib/pq" //なんかよくわからんけどいる　これがないとDBアクセスがruntimeErrorを吐く
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

// SetConnect 接続を取得する
// 毎回作らなくてもメモリ上のどこかに保存しておけないのか？
func GetConfig() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     os.Getenv("GoogleClientID"),
		ClientSecret: os.Getenv("googleClientSecret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("AuthorizeEndpoint"),
			TokenURL: os.Getenv("TokenEndpoint"),
		},
		Scopes:      []string{"https://www.googleapis.com/auth/userinfo.email"},
		RedirectURL: "http://localhost:8080/google/callback",
	}

	return config
}

func Auth(c *gin.Context) {
	conf := GetConfig()
	state := `hoge`

	// stateをsessionなどに保存.

	// リダイレクトURL作成.
	redirectURL := conf.AuthCodeURL(state)

	//ターミナルに吐くとおかしなエンコーディングになるのでエンコーディングを戻す
	fmt.Println(redirectURL)

	// redirectURLをクライアントに返す.
	c.JSON(200, gin.H{
		"redirectUrl": redirectURL,
	})
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
	r.GET("/auth", Auth)
	r.GET("/google/callback", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login completed",
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
	webServerTest()
}
