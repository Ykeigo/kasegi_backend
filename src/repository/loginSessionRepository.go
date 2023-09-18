package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	models "kasegi/my_models"
	util "kasegi/util"

	"github.com/volatiletech/sqlboiler/v4/boil"

	_ "github.com/lib/pq" //なんかよくわからんけどいる　これがないとDBアクセスがruntimeErrorを吐く
)

type LoginSessionRepository struct{
}

func (ur LoginSessionRepository)Insert(loginSession models.LoginSession, db *sql.DB) {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568
	loginSession.ID = util.IdGenerator{}.GenerateSurrogateKey()
	loginSession.ExpiresAt = time.Now().Add(time.Hour * 24 * 30)
	err := loginSession.Insert(context.Background(), db, boil.Infer())
	fmt.Println(err)
}

func (ur LoginSessionRepository)Get(db *sql.DB) models.LoginSession {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	loginSession, e := models.FindLoginSession(context.Background(), db, "")
	fmt.Println(e)
	return *loginSession
}

func (ur LoginSessionRepository)List(db *sql.DB) []models.LoginSession {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	//全部とる
	loginSessionSlice, e := models.LoginSessions().All(context.Background(), db)
	fmt.Println(e)

	var loginSession []models.LoginSession
	for _, l := range loginSessionSlice {
		loginSession = append(loginSession, *l)
	}

	return loginSession
}

func (ur LoginSessionRepository)DeleteByUserId(userId string, db *sql.DB) []models.LoginSession {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	//全部とる
	loginSessionSlice, e := models.LoginSessions(models.LoginSessionWhere.UserID.EQ(userId)).All(context.Background(), db)
	fmt.Println(e)

	var loginSession []models.LoginSession
	for _, l := range loginSessionSlice {
		l.Delete(context.Background(), db)
	}

	return loginSession
}

func (ur LoginSessionRepository)FindBySessionToken(sessionToken string, db *sql.DB) []models.LoginSession {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	//全部とる

	loginSessionSlice, e := models.LoginSessions(models.LoginSessionWhere.SessionToken.EQ(sessionToken)).All(context.Background(), db)
	fmt.Println(e)

	var loginSession []models.LoginSession
	for _, l := range loginSessionSlice {
		loginSession = append(loginSession, *l)
	}

	return loginSession
}