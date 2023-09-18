package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "kasegi/my_models"
	util "kasegi/util"

	"github.com/volatiletech/sqlboiler/v4/boil"

	_ "github.com/lib/pq" //なんかよくわからんけどいる　これがないとDBアクセスがruntimeErrorを吐く
)

type UserRepository struct{
}

func (ur UserRepository)Insert(user models.User, db *sql.DB) {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568
	user.ID = util.IdGenerator{}.GenerateSurrogateKey()
	err := user.Insert(context.Background(), db, boil.Infer())
	fmt.Println(err)
}

func (ur UserRepository)Get(db *sql.DB) models.User {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	user, e := models.FindUser(context.Background(), db, "")
	fmt.Println(e)
	return *user
}

func (ur UserRepository)List(db *sql.DB) []models.User {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	//全部とる
	usersSlice, e := models.Users().All(context.Background(), db)
	fmt.Println(e)

	var users []models.User
	for _, u := range usersSlice {
		users = append(users, *u)
	}

	return users
}

func (ur UserRepository)FindByEmail(email string, db *sql.DB) []models.User {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	//全部とる

	usersSlice, e := models.Users(models.UserWhere.Email.EQ(email)).All(context.Background(), db)
	fmt.Println(e)

	var users []models.User
	for _, u := range usersSlice {
		users = append(users, *u)
	}

	return users
}