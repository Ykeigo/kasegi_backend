package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "kasegi/my_models"
	util "kasegi/util"

	_ "github.com/lib/pq" //なんかよくわからんけどいる　これがないとDBアクセスがruntimeErrorを吐く
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ChecklistTemplateRepository struct{

}

type TemplateCheckItem struct {
	Title string
	IsChecked bool
}

type ChecklistTemplate struct {
	Id string
	GameTitleId string
	CreatedByUserId string
	CheckItems []TemplateCheckItem
}

func (ur ChecklistTemplateRepository)Insert(checklistTemplate ChecklistTemplate, db *sql.DB) {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568
	var checklistTemplateDbModel = models.ChecklistTemplate{
		ID: util.IdGenerator{}.GenerateSurrogateKey(),
		GameTitleID: checklistTemplate.GameTitleId,
		CreatedByUserID: checklistTemplate.CreatedByUserId,
	}

	for _, checkItem := range checklistTemplate.CheckItems {
		var ChecklistTemplateItemsDbModel = models.ChecklistTemplateItem{
			ID: util.IdGenerator{}.GenerateSurrogateKey(),
			ChecklistTemplateID: checklistTemplateDbModel.ID,
			Title: checkItem.Title,
			IsChecked: checkItem.IsChecked,
		}
		ChecklistTemplateItemsDbModel.Insert(context.Background(), db, boil.Infer())
	}
	checklistTemplateDbModel.Insert(context.Background(), db, boil.Infer())
}

func convertChecklistTemplateDbModelToChecklistTemplate(checklistTemplateDbModel models.ChecklistTemplate, ChecklistTemplateItemsDbModels []*models.ChecklistTemplateItem) ChecklistTemplate {
	var checkItems []TemplateCheckItem
	for _, ChecklistTemplateItemsDbModel := range ChecklistTemplateItemsDbModels {
		var checkItem = TemplateCheckItem{
			Title: ChecklistTemplateItemsDbModel.Title,
			IsChecked: ChecklistTemplateItemsDbModel.IsChecked,
		}
		checkItems = append(checkItems, checkItem)
	}
	var checklistTemplate = ChecklistTemplate{
		Id: checklistTemplateDbModel.ID,
		GameTitleId: checklistTemplateDbModel.GameTitleID,
		CreatedByUserId: checklistTemplateDbModel.CreatedByUserID,
		CheckItems: checkItems,	
	}
	return checklistTemplate
}

func (ur ChecklistTemplateRepository)Get(id string, db *sql.DB) ChecklistTemplate {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	checklistTemplate, e1 := models.FindChecklistTemplate(context.Background(), db, id)
	fmt.Println(e1)

	checklistTemplateSelectionItems, e2 := models.ChecklistTemplateItems(models.ChecklistTemplateItemWhere.ChecklistTemplateID.EQ(checklistTemplate.ID)).All(context.Background(), db)
	fmt.Println(e2)

	return convertChecklistTemplateDbModelToChecklistTemplate(*checklistTemplate, checklistTemplateSelectionItems)
}

func (ur ChecklistTemplateRepository)List(db *sql.DB) []ChecklistTemplate {
	checklistTemplateDbModels, e1 := models.ChecklistTemplates().All(context.Background(), db)
	fmt.Println(e1)

	checklistTemplateSelectionItemsDbModels, e2 := models.ChecklistTemplateItems().All(context.Background(), db)
	fmt.Println(e2)

	var checklistTemplates []ChecklistTemplate
	for _, checklistTemplate := range checklistTemplateDbModels {
		var filterdItems []*models.ChecklistTemplateItem
		for _, checklistTemplateSelectionItem := range checklistTemplateSelectionItemsDbModels {
			if checklistTemplateSelectionItem.ChecklistTemplateID == checklistTemplate.ID {
				filterdItems = append(filterdItems, checklistTemplateSelectionItem)
			}
		}

		var a = convertChecklistTemplateDbModelToChecklistTemplate(*checklistTemplate, filterdItems)
		checklistTemplates = append(checklistTemplates, a)
	}

	return checklistTemplates
}

func (ur ChecklistTemplateRepository)FindByUserId(userId string, db *sql.DB) []ChecklistTemplate {
	// https://qiita.com/hiro9/items/e6e41ec822a7077c3568

	//全部とる

	checklistTemplateesSlice, e := models.ChecklistTemplates(models.ChecklistTemplateWhere.CreatedByUserID.EQ(userId)).All(context.Background(), db)
	fmt.Println(e)
	
	checklistTemplateSelectionItemsDbModels, e2 := models.ChecklistTemplateItems().All(context.Background(), db)
	fmt.Println(e2)

	
	var checklistTemplates []ChecklistTemplate
	for _, checklistTemplate := range checklistTemplateesSlice {
		var filterdItems []*models.ChecklistTemplateItem
		for _, checklistTemplateSelectionItem := range checklistTemplateSelectionItemsDbModels {
			if checklistTemplateSelectionItem.ChecklistTemplateID == checklistTemplate.ID {
				filterdItems = append(filterdItems, checklistTemplateSelectionItem)
			}
		}

		var a = convertChecklistTemplateDbModelToChecklistTemplate(*checklistTemplate, filterdItems)
		checklistTemplates = append(checklistTemplates, a)
	}

	return checklistTemplates
}