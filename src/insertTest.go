package main

import (
	"context"
	"database/sql"
	"fmt"

	models "test-app/models_psql"

	_ "github.com/lib/pq" //なんかよくわからんけどいる　ないとruntimeErrorを吐く
)

func main() {
	db, _ := sql.Open(
		"postgres",
		"host=postgres dbname=example user=postgres password=password sslmode=disable")
	titles, _ := models.GameTitles().All(context.Background(), db)
	fmt.Println("jets:", titles)
}
