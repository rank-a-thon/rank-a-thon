package controllers

import (
	"github.com/gin-gonic/gin"
)

type JudgeController struct{}



func (ctrl RankerController) GetEvaluations(context *gin.Context) {
	// get list of assigned evaluations for a judge
	if userID := getUserID(context); userID != 0 {

	}
}

func (ctrl RankerController) GetEvaluation(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

	}
}

func (ctrl RankerController) SetEvaluation(context *gin.Context) {
	if userID := getUserID(context); userID != 0 {

	}
}
