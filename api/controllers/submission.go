package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"
	"log"
	"net/http"
	"os"
	"strconv"
)

type SubmissionController struct{}

var submissionModel = new(models.SubmissionModel)
var submissionLikeModel = new(models.SubmissionLikeModel)

// Create submission
func (ctrl SubmissionController) Create(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		var submissionForm forms.SubmissionForm

		if context.ShouldBindJSON(&submissionForm) != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			context.Abort()
			return
		}
		teamID, err := getTeamIDForEvent(context, userID)
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

// Get all submissions belonging to user ID
func (ctrl SubmissionController) AllForUserID(context *gin.Context) {
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

// Get all submissions belonging to event
func (ctrl SubmissionController) AllForEvent(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		event, err := fetchAndValidateEvent(context)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid event name"})
			context.Abort()
			return
		}

		data, err := submissionModel.AllForEvent(models.Event(event))
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get submissions", "error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// get submission for event from userID
func (ctrl SubmissionController) One(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		teamID, err := getTeamIDForEvent(context, userID)
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

// Update submission
func (ctrl SubmissionController) Update(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		teamID, err := getTeamIDForEvent(context, userID)
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

// Delete Submission
func (ctrl SubmissionController) Delete(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		teamID, err := getTeamIDForEvent(context, userID)
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

// Like submission
func (ctrl SubmissionController) LikeSubmission(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		submissionID, err := strconv.ParseUint(context.Query("submission-id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
			context.Abort()
			return
		}

		_, err = submissionLikeModel.Create(uint(submissionID), userID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Submission could not be liked", "error": err.Error()})
			context.Abort()
			return
		}

		err = submissionModel.IncrementLike(uint(submissionID), 1)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Submission could not be liked", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Submission liked"})
	}
}

// Unlike submission
func (ctrl SubmissionController) UnlikeSubmission(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		submissionID, err := strconv.ParseUint(context.Query("submission-id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
			context.Abort()
			return
		}

		err = submissionLikeModel.Delete(uint(submissionID), userID)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Submission could not be unliked", "error": err.Error()})
			context.Abort()
			return
		}

		err = submissionModel.IncrementLike(uint(submissionID), -1)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Submission could not be unliked", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Submission unliked"})
	}
}

// Upload file
func (ctrl SubmissionController) UploadFile(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {
		teamID, err := getTeamIDForEvent(context, userID)

		event, err := fetchAndValidateEvent(context)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid event name"})
			context.Abort()
			return
		}

		file, err := context.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}

		imageFolder := fmt.Sprintf("submission_files/%s/%d", event, teamID)
		imagePath := fmt.Sprintf("submission_files/%s/%d/%s", event, teamID, file.Filename)
		if _, err := os.Stat(imageFolder); os.IsNotExist(err) {
			err = os.MkdirAll(imageFolder, 0777)
			if err != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"message": "Error creating directory", "error": err.Error()})
				context.Abort()
				return
			}
		}

		err = context.SaveUploadedFile(file, imagePath)
		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Error uploading Image", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"image_url": "api/v1/" + imagePath})
	}
}