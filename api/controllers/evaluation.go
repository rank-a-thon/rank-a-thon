package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"
)

type EvaluationController struct{}

var evaluationModel = new(models.EvaluationModel)

func (ctrl EvaluationController) All(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		isJudge, err := userModel.IsJudgeForUserID(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get evaluations", "error": err.Error()})
			context.Abort()
			return
		}
		if !isJudge {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not get evaluations, user is not a Judge"})
			context.Abort()
			return
		}

		data, err := evaluationModel.All(userID)

		if err != nil {
			context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get evaluations", "error": err.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func (ctrl EvaluationController) One(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		isJudge, err := userModel.IsJudgeForUserID(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get evaluation", "error": err.Error()})
			context.Abort()
			return
		}
		if !isJudge {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not get evaluation, user is not a Judge"})
			context.Abort()
			return
		}

		id := context.Param("id")
		if id, err := strconv.ParseUint(id, 10, 64); err == nil {

			data, err := evaluationModel.One(uint(id))

			if err != nil {
				context.JSON(http.StatusNotFound, gin.H{"Message": "Evaluation not found", "error": err.Error()})
				context.Abort()
				return
			}

			context.JSON(http.StatusOK, gin.H{"data": data})

		} else {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		}
	}
}

func (ctrl EvaluationController) Update(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

		isJudge, err := userModel.IsJudgeForUserID(userID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get evaluation", "error": err.Error()})
			context.Abort()
			return
		}
		if !isJudge {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not get evaluation, user is not a Judge"})
			context.Abort()
			return
		}

		id := context.Param("id")
		if id, err := strconv.ParseUint(id, 10, 64); err == nil {

			var evaluationForm forms.EvaluationForm

			if context.ShouldBindJSON(&evaluationForm) != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
				context.Abort()
				return
			}

			err := evaluationModel.Update(uint(id), evaluationForm)
			if err != nil {
				context.JSON(http.StatusNotAcceptable, gin.H{"Message": "Evaluation could not be updated", "error": err.Error()})
				context.Abort()
				return
			}

			context.JSON(http.StatusOK, gin.H{"message": "Evaluation updated"})

		} else {
			context.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter", "error": err.Error()})
		}
	}
}
