package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"
)

type SubmissionController struct{}

var submissionModel = new(models.SubmissionModel)

func (ctrl SubmissionController) Create(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		var submissionForm forms.SubmissionForm

		if context.ShouldBindJSON(&submissionForm) != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			context.Abort()
			return
		}
		teamID, err := ctrl.getTeamIDForEvent(context, userID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Could not fetch team ID", "error": err.Error()})
			context.Abort()
			return
		}
		submissionID, err := submissionModel.Create(teamID, submissionForm)

		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Submission could not be created", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Submission created", "id": submissionID})
	}
}

func (ctrl SubmissionController) All(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		data, err := submissionModel.AllForUserID(userID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get submissions", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func (ctrl SubmissionController) One(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		teamID, err := ctrl.getTeamIDForEvent(context, userID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Could not fetch team ID", "error": err.Error()})
			context.Abort()
			return
		}
		data, err := submissionModel.OneByTeamID(teamID)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Submission not found", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": data})

	}
}

func (ctrl SubmissionController) Update(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		teamID, err := ctrl.getTeamIDForEvent(context, userID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Could not fetch team ID", "error": err.Error()})
			context.Abort()
			return
		}

		var submissionForm forms.SubmissionForm
		if context.ShouldBindJSON(&submissionForm) != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			context.Abort()
			return
		}

		err = submissionModel.Update(teamID, submissionForm)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Submission could not be updated", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Submission updated"})
	}
}

func (ctrl SubmissionController) Delete(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		teamID, err := ctrl.getTeamIDForEvent(context, userID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Could not fetch team ID", "error": err.Error()})
			context.Abort()
			return
		}
		err = submissionModel.Delete(teamID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Submission could not be deleted", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Submission deleted"})
	}
}

func (ctrl SubmissionController) getTeamIDForEvent(context *gin.Context, userID uint) (teamID uint, err error){
	event, err := fetchAndValidateEvent(context)
	if err != nil {
		return 0, errors.New("invalid event name")
	}

	user, err := userModel.One(userID)
	if err != nil {
		return 0, errors.New("unable to fetch user")
	}
	teamIDForEvent := userModel.GetTeamIDForEventMap(user)
	return teamIDForEvent[event], nil
}
