package api

import (
	"database/sql"
	repository "kasegi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameMatchCrudApi struct{
}

type CreateGameMatchRequest struct {
	SessionToken string
	GameMatch repository.GameMatch
}

func (gmca GameMatchCrudApi)CreateGameMatch(c *gin.Context, db *sql.DB) {
	var loginSessionRepository repository.LoginSessionRepository
	var gameMatchRepository repository.GameMatchRepository

	var request CreateGameMatchRequest
	
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var sessions = loginSessionRepository.FindBySessionToken(request.SessionToken, db)
	if len(sessions) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session", "session" : request.SessionToken})
		return
	}

	gameMatchRepository.Insert(request.GameMatch, db)
	
	c.JSON(200, gin.H{
		"message": "match created.",
		"gameMatch": request.GameMatch,
	})
	return
}