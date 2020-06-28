package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"
)

type TeamController struct{}

var teamModel = new(models.TeamModel)
var teamInviteModel = new(models.TeamInviteModel)

func (ctrl TeamController) Create(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		var teamForm forms.TeamForm

		if context.ShouldBindJSON(&teamForm) != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			context.Abort()
			return
		}

		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		} else if user.TeamID != 0 {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Team could not be created, user already has team"})
			context.Abort()
			return
		}

		teamID, err := teamModel.Create(userID, teamForm)

		if teamID == 0 && err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Team could not be created", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Team created", "id": teamID})
	}
}

func (ctrl TeamController) One(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		} else if user.TeamID == 0 {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "User does not have team"})
			context.Abort()
			return
		}

		data, err := teamModel.One(user.TeamID)

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Team not found", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func (ctrl TeamController) Update(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		} else if user.TeamID == 0 {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "User does not have team"})
			context.Abort()
			return
		}

		var teamForm forms.TeamForm

		if context.ShouldBindJSON(&teamForm) != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			context.Abort()
			return
		}

		err = teamModel.Update(user.TeamID, teamForm)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Team could not be updated", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Team updated"})
	}
}

//TODO
//send invite
//accept invite
//decline invite

func (ctrl TeamController) Delete(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		user, err := userModel.One(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user", "error": err.Error()})
			context.Abort()
			return
		} else if user.TeamID == 0 {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "User does not have team"})
			context.Abort()
			return
		}

		err = teamModel.Delete(user.TeamID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Team could not be deleted", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
	}
}
