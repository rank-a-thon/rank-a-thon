package controllers

import (
	"github.com/gin-gonic/gin"
)

type RankerController struct{}

func (ctrl RankerController) CreateEvaluations(context *gin.Context) {
	// get list of judges, submissions and create evaluations
	if userID := getUserID(context); userID != 0 {

	}
}

func (ctrl RankerController) GetRankings(context *gin.Context) {
	// check if all judges have finished evaluation, if true return teams in sorted ranking
	if userID := getUserID(context); userID != 0 {

	}
}

