package main

import (
	"context"
	"database/sql"
	"fmt"

	models "test-app/models_psql"

	_ "github.com/lib/pq" //なんかよくわからんけどいる　ないとruntimeErrorを吐く
)

func main() {
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
