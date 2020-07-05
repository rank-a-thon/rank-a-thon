package controllers

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rank-a-thon/rank-a-thon/api/models"
)

type RankerController struct{}

func (ctrl RankerController) CreateEvaluations(context *gin.Context) {
	// get list of judges, submissions and create evaluations
	userID := getUserID(context);
	if userID != 3 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization"})
		context.Abort()
		return
	}

	judges, err := userModel.GetAllJudges()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching judges", "error": err.Error()})
	}
	event, err := fetchAndValidateEvent(context)
	submissions, err := submissionModel.AllForEvent(models.Event(event))

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(submissions), func(i, j int) { submissions[i], submissions[j] = submissions[j], submissions[i] })
	submissionsPerJudge := int(math.Floor(float64(len(submissions)) / float64(len(judges))))
	startIdx, endIdx := 0, 0
	for _, judge := range judges {
		endIdx = startIdx + submissionsPerJudge
		if endIdx > len(submissions) {
			endIdx = len(submissions)
		}
		for i := startIdx; i < endIdx; i++ {
			_, err := evaluationModel.Create(judge.ID, submissions[i].ID)
			if err != nil {
				log.Println(err)
			}
		}
		startIdx += submissionsPerJudge
	}

	for i := 0; startIdx < len(submissions); i, startIdx = i+1, startIdx+1 { // distribute the remainder of submissions
		_, err := evaluationModel.Create(judges[i].ID, submissions[startIdx].ID)
		if err != nil {
			log.Println(err)
		}
	}
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Created evaluations with error", "error": err.Error()})
		context.Abort()
		return
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Created evaluations"})
	}

}

func normaliseJudgeScores(context *gin.Context) {

}

func (ctrl RankerController) CalculateTeamRankings(context *gin.Context) {
	/*
	normaliseJudgeScore(context)
	For each team in event
	get all evaluations assigned to submission
	calculate mean of normalised mean for each category
	Save into psql
	 */
}

func (ctrl RankerController) GetTeamRankings(context *gin.Context) {
	// check if all judges have finished evaluation, if true return teams in sorted ranking
	if userID := getUserID(context); userID == 3 {

	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization"})
	}
}

