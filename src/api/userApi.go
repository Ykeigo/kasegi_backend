package api

import (
	"database/sql"
	repository "kasegi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct{
}

type GetMyUserRequest struct {
	SessionToken string
}

func (ua UserApi)GetMyUser(c *gin.Context, db *sql.DB) {
	var loginSessionRepository repository.LoginSessionRepository
	var userRepository repository.UserRepository

	var request GetMyUserRequest
	
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var sessions = loginSessionRepository.FindBySessionToken(request.SessionToken, db)
	if len(sessions) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session", "session" : request.SessionToken})
		return
	}

	var userId = sessions[0].UserID
	var myUser = userRepository.FindByUserId(userId, db)
	if len(myUser) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found", "session" : sessions[0]})
		return
	}else{	
		c.JSON(200, gin.H{
			"message": "match created.",
			"user": myUser,
		})
		return
	}
}