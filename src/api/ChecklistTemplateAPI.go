package api

import (
	"database/sql"
	repository "kasegi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChecklistTemplateApi struct{
}

type CreateChecklistTemplateRequest struct {
	SessionToken string
	ChecklistTemplate repository.ChecklistTemplate
}

type GetMyChecklistTemplateRequest struct {
	SessionToken string
}

func (gmca ChecklistTemplateApi)CreateMyChecklistTempalte(c *gin.Context, db *sql.DB) {
	var loginSessionRepository repository.LoginSessionRepository
	var checklistTemplateRepository repository.ChecklistTemplateRepository

	var request CreateChecklistTemplateRequest
	
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var sessions = loginSessionRepository.FindBySessionToken(request.SessionToken, db)
	if len(sessions) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session", "session" : request.SessionToken})
		return
	}

	request.ChecklistTemplate.CreatedByUserId = sessions[0].UserID
	checklistTemplateRepository.Insert(request.ChecklistTemplate, db)
	
	c.JSON(200, gin.H{
		"message": "checklist template created.",
		"gameMatch": request.ChecklistTemplate,
	})
	return
}

func (gmca ChecklistTemplateApi)ListMyChecklistTemplate(c *gin.Context, db *sql.DB) {
	var loginSessionRepository repository.LoginSessionRepository
	var checklistTemplateRepository repository.ChecklistTemplateRepository

	var request GetMyChecklistTemplateRequest
	
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var sessions = loginSessionRepository.FindBySessionToken(request.SessionToken, db)
	if len(sessions) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session", "session" : request.SessionToken})
		return
	}
	var templates = checklistTemplateRepository.FindByUserId(sessions[0].UserID, db)

	
	c.JSON(200, gin.H{
		"checklistTemplate": templates,
	})
	return
}