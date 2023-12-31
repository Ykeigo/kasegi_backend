package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	models "kasegi/my_models"
	util "kasegi/util"

	_ "github.com/lib/pq" //なんかよくわからんけどいる　これがないとDBアクセスがruntimeErrorを吐く
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type GameMatchRepository struct{

}

type CheckItem struct {
	Title string
	IsChecked bool
}

type GameMatch struct {
	Id string
	GameId string
	UserId string
	CheckItems []CheckItem
	CreatedAt time.Time
}

func (ur GameMatchRepository)Insert(gameMatch GameMatch, db *sql.DB) {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568
	var gameMatchDbModel = models.GameMatch{
		ID: util.IdGenerator{}.GenerateSurrogateKey(),
		GameID: gameMatch.GameId,
		UserID: gameMatch.UserId,
		CreatedAt: time.Now(),
	}

	for _, checkItem := range gameMatch.CheckItems {
		var gameMatchCheckItemDbModel = models.GameMatchChecklistItem{
			ID: util.IdGenerator{}.GenerateSurrogateKey(),
			MatchID: gameMatchDbModel.ID,
			Title: checkItem.Title,
			IsChecked: checkItem.IsChecked,
		}
		gameMatchCheckItemDbModel.Insert(context.Background(), db, boil.Infer())
	}
	gameMatchDbModel.Insert(context.Background(), db, boil.Infer())
}

func convertGameMatchDbModelToGameMatch(gameMatchDbModel models.GameMatch, gameMatchCheckItemDbModels []*models.GameMatchChecklistItem) GameMatch {
	var checkItems []CheckItem
	for _, gameMatchCheckItemDbModel := range gameMatchCheckItemDbModels {
		var checkItem = CheckItem{
			Title: gameMatchCheckItemDbModel.Title,
			IsChecked: gameMatchCheckItemDbModel.IsChecked,
		}
		checkItems = append(checkItems, checkItem)
	}
	var gameMatch = GameMatch{
		Id: gameMatchDbModel.ID,
		GameId: gameMatchDbModel.GameID,
		UserId: gameMatchDbModel.UserID,
		CheckItems: checkItems,
		CreatedAt: gameMatchDbModel.CreatedAt,
	}
	return gameMatch
}

func (ur GameMatchRepository)Get(id string, db *sql.DB) GameMatch {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	gameMatch, e1 := models.FindGameMatch(context.Background(), db, id)
	fmt.Println(e1)

	gameMatchSelectionItems, e2 := models.GameMatchChecklistItems(models.GameMatchChecklistItemWhere.MatchID.EQ(gameMatch.ID)).All(context.Background(), db)
	fmt.Println(e2)

	return convertGameMatchDbModelToGameMatch(*gameMatch, gameMatchSelectionItems)
}

func (ur GameMatchRepository)List(db *sql.DB) []GameMatch {
	gameMatchDbModels, e1 := models.GameMatches().All(context.Background(), db)
	fmt.Println(e1)

	gameMatchSelectionItemsDbModels, e2 := models.GameMatchChecklistItems().All(context.Background(), db)
	fmt.Println(e2)

	var gameMatchs []GameMatch
	for _, gameMatch := range gameMatchDbModels {
		var filterdItems []*models.GameMatchChecklistItem
		for _, gameMatchSelectionItem := range gameMatchSelectionItemsDbModels {
			if gameMatchSelectionItem.MatchID == gameMatch.ID {
				filterdItems = append(filterdItems, gameMatchSelectionItem)
			}
		}

		var a = convertGameMatchDbModelToGameMatch(*gameMatch, filterdItems)
		gameMatchs = append(gameMatchs, a)
	}

	return gameMatchs
}

func (ur GameMatchRepository)FindByUserId(userId string, db *sql.DB) []GameMatch {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	//全部とる

	gameMatchesSlice, e := models.GameMatches(models.GameMatchWhere.UserID.EQ(userId)).All(context.Background(), db)
	fmt.Println(e)
	
	gameMatchSelectionItemsDbModels, e2 := models.GameMatchChecklistItems().All(context.Background(), db)
	fmt.Println(e2)

	
	var gameMatchs []GameMatch
	for _, gameMatch := range gameMatchesSlice {
		var filterdItems []*models.GameMatchChecklistItem
		for _, gameMatchSelectionItem := range gameMatchSelectionItemsDbModels {
			if gameMatchSelectionItem.MatchID == gameMatch.ID {
				filterdItems = append(filterdItems, gameMatchSelectionItem)
			}
		}

		var a = convertGameMatchDbModelToGameMatch(*gameMatch, filterdItems)
		gameMatchs = append(gameMatchs, a)
	}

	return gameMatchs
}