package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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

		submissionID, err := submissionModel.Create(userID, submissionForm)

		if submissionID == 0 && err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"message": "Submission could not be created", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Submission created", "id": submissionID})
	}
}

func (ctrl SubmissionController) All(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		data, err := submissionModel.All(userID)

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

		id := context.Param("id")
		if id, err := strconv.ParseUint(id, 10, 64); err == nil {

			data, err := submissionModel.One(userID, uint(id))

			if err != nil {
				context.JSON(http.StatusNotFound, gin.H{"Message": "Submission not found", "error": err.Error()})
				context.Abort()
				return
			}

			context.JSON(http.StatusOK, gin.H{"data": data})

		} else {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		}
	}
}

func (ctrl SubmissionController) Update(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		id := context.Param("id")
		if id, err := strconv.ParseUint(id, 10, 64); err == nil {

			var submissionForm forms.SubmissionForm

			if context.ShouldBindJSON(&submissionForm) != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
				context.Abort()
				return
			}

			err := submissionModel.Update(userID, uint(id), submissionForm)
			if err != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Submission could not be updated", "error": err.Error()})
				context.Abort()
				return
			}

			context.JSON(http.StatusOK, gin.H{"message": "Submission updated"})

		} else {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter", "error": err.Error()})
		}
	}
}

func (ctrl SubmissionController) Delete(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		id := context.Param("id")
		if id, err := strconv.ParseUint(id, 10, 64); err == nil {

			err := submissionModel.Delete(userID, uint(id))
			if err != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Submission could not be deleted", "error": err.Error()})
				context.Abort()
				return
			}

			context.JSON(http.StatusOK, gin.H{"message": "Submission deleted"})

		} else {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		}
	}
}
