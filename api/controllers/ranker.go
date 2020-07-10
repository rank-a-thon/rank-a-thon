package controllers

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rank-a-thon/rank-a-thon/api/forms"
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

func (ctrl RankerController) CalculateTeamRankings(context *gin.Context) {
	/*
	normaliseJudgeScores(context, event)
	For each team in event
	get all evaluations assigned to submission
	calculate mean of normalised mean for each category
	Save into psql
	 */
	userID := getUserID(context);
	if userID != 3 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization"})
		context.Abort()
		return
	}
	event, err := fetchAndValidateEvent(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid event name"})
		context.Abort()
		return
	}

	normaliseJudgeScores(context, event)
	// get all normalised scores for each submission and take mean
	// rank submissions by each category and persist
}

func normaliseJudgeScores(context *gin.Context, event string) {
	/*
		For each judge
		Get mean and std of scores assign for each category
		Normalise by each category
		Save a new copy with normalised = true
	*/
	judges, err := userModel.GetAllJudges()	// TODO in future need to get all judges for event
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching judges", "error": err.Error()})
	}
	for _, judge := range judges {
		evaluations, err := evaluationModel.AllForJudge(judge.ID)
		if err != nil {
			log.Println("Error fetching evaluations for judge: " + err.Error())
		}

		mean, std := calculateStatisticsForEvaluations(evaluations)

		for _, evaluation := range evaluations {
			standardisedRatingArray := evaluation.ReadRatingsIntoArray()
			for i := 0; i < models.NumberOfRatings; i++ {
				standardisedRatingArray[i] = (standardisedRatingArray[i] - mean[i]) / std[i]
			}
			form := forms.EvaluationForm{
				MainRating:             standardisedRatingArray[0],
				AnnoyingRating:         standardisedRatingArray[1],
				EntertainRating:        standardisedRatingArray[2],
				BeautifulRating:        standardisedRatingArray[3],
				SociallyUsefulRating:   standardisedRatingArray[4],
				HardwareRating:         standardisedRatingArray[5],
				AwesomelyUselessRating: standardisedRatingArray[6],
			}
			_, err := evaluationModel.CreateStandardised(judge.ID, evaluation.SubmissionID, form)

			if err != nil {
				log.Println(err.Error())
			}
		}

	}

}

func calculateStatisticsForEvaluations(evaluations []models.Evaluation) (mean []float64, std []float64) {
	sum := make([]float64, models.NumberOfRatings)
	for _, evaluation := range evaluations {
		ratingArray := evaluation.ReadRatingsIntoArray()
		for i := 0; i < models.NumberOfRatings; i++ {
			sum[i] += ratingArray[i]
		}
	}

	mean = make([]float64, models.NumberOfRatings)
	for i := 0; i < 7; i++ {
		mean[i] = sum[i] / float64(len(evaluations))
	}

	std = make([]float64, models.NumberOfRatings)
	for _, evaluation := range evaluations {
		ratingArray := evaluation.ReadRatingsIntoArray()
		for i := 0; i < models.NumberOfRatings; i++ {
			std[i] += math.Pow(ratingArray[i] - mean[i], 2)
		}
	}

	for i := 0; i < models.NumberOfRatings; i++ {
		std[i] = math.Sqrt(std[i] / float64(len(evaluations)))
	}
	return mean, std
}

func (ctrl RankerController) GetTeamRankings(context *gin.Context) {
	// check if all judges have finished evaluation, if true return teams in sorted ranking
	userID := getUserID(context);
	if userID != 3 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization"})
		context.Abort()
		return
	}
	_, err := fetchAndValidateEvent(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid event name"})
		context.Abort()
		return
	}

}

