package controllers

import (
	"fmt"
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

var submissionRankingModel = new(models.SubmissionRankingModel)

// Once all the teams have made their submissions, create evaluations for each judge-submission pair
func (ctrl RankerController) CreateEvaluations(context *gin.Context) {
	// get list of judges, submissions and create evaluations
	userID := getUserID(context)
	user, err := userModel.One(userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching user", "error": err.Error()})
		context.Abort()
		return
	}

	if user.UserType != 3 {
		fmt.Println(user.UserType)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization"})
		context.Abort()
		return
	}

	judges, err := userModel.GetAllJudges()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching judges", "error": err.Error()})
		context.Abort()
		return
	}

	event, err := fetchAndValidateEvent(context)
	submissions, err := submissionModel.AllForEvent(models.Event(event))

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(submissions), func(i, j int) { submissions[i], submissions[j] = submissions[j], submissions[i] })

	submissionsPerJudge := int(math.Floor(float64(len(submissions)) / float64(len(judges))))
	fmt.Println(float64(len(submissions)))
	fmt.Println(float64(len(judges)))
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
		context.JSON(http.StatusOK, gin.H{"message": "Created evaluations"})
	}

}

// Calculate team rankings and persist
func (ctrl RankerController) CalculateTeamRankings(context *gin.Context) {
	/*
		normaliseJudgeScores(context, event)
		For each team in event
		get all evaluations assigned to submission
		calculate mean of normalised mean for each category
		Save into psql
	*/
	userID := getUserID(context)
	user, err := userModel.One(userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching user", "error": err.Error()})
		context.Abort()
		return
	}

	if user.UserType != 3 {
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
	// get all normalised scores for each submission, take mean and persist to submissionRankings
	allSubmissions, err := submissionModel.AllForEvent(models.Event(event))
	for _, submission := range allSubmissions {
		evaluationsForSubmission, err := evaluationModel.AllForSubmission(submission.ID)
		if err != nil {
			log.Println(err.Error())
		}
		var mean = make([]float64, models.NumberOfRatings)
		var count int = 0
		for _, evaluation := range evaluationsForSubmission {
			if evaluation.Normalised {
				ratingArray := evaluation.ReadRatingsIntoArray()
				for i := 0; i < models.NumberOfRatings; i++ {
					mean[i] += ratingArray[i]
				}
				count++
			}
		}

		for i := 0; i < models.NumberOfRatings; i++ {
			mean[i] /= float64(count)
		}

		_, err = submissionRankingModel.Create(submission.ID, mean)
		if err != nil {
			log.Println(err)
		}
	}
	// for each ranking type, sort by that value and persist each of the rankings
	categories := []string{"main", "annoying", "entertaining", "beautiful", "socially_useful", "hardware", "awesomely_useless"}

	for _, category := range categories {
		rankedSubmissions, err := submissionRankingModel.AllByCategory(category)
		if err != nil {
			log.Println(err)
		}

		for ranking, submissionRanking := range rankedSubmissions {
			err = submissionRankingModel.Update(submissionRanking.SubmissionID, map[string]uint{
				fmt.Sprintf("%s_ranking", category): uint(ranking),
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
	context.JSON(http.StatusOK, gin.H{"message": "Teams ranked successfully"})

}

// Normalise a judge's scores based on each judging category
func normaliseJudgeScores(context *gin.Context, event string) {
	/*
		For each judge
		Get mean and std of scores assign for each category
		Normalise by each category
		Save a new copy with normalised = true
	*/
	judges, err := userModel.GetAllJudges() // TODO in future need to get all judges for event
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
				if std[i] != 0 {
					standardisedRatingArray[i] = (standardisedRatingArray[i] - mean[i]) / std[i]
				} else {
					standardisedRatingArray[i] = standardisedRatingArray[i] - mean[i]
				}
			}

			form := forms.EvaluationFormFloat{
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

// Calculate mean and std of scores
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
			std[i] += math.Pow(ratingArray[i]-mean[i], 2)
		}
	}

	for i := 0; i < models.NumberOfRatings; i++ {
		std[i] = math.Sqrt(std[i] / float64(len(evaluations)))
	}
	return mean, std
}

// Get team rankings by range from inclusive to exclusive
func (ctrl RankerController) GetTeamRankingsByRange(context *gin.Context) {
	// check if all judges have finished evaluation, if true return teams in sorted ranking
	userID := getUserID(context)
	user, err := userModel.One(userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error fetching user", "error": err.Error()})
		context.Abort()
		return
	}

	if user.UserType != 3 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization"})
		context.Abort()
		return
	}

	_, err = fetchAndValidateEvent(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid event name"})
		context.Abort()
		return
	}

	var rankerForm forms.RankerForm
	err = context.ShouldBindJSON(&rankerForm)
	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form", "error": err.Error()})
		context.Abort()
		return
	}

	submissionRankings, err := submissionRankingModel.AllByCategory(rankerForm.Category)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"Message": "Error fetching submission rankings", "error": err.Error()})
		context.Abort()
		return
	}

	selectedRankings := submissionRankings[rankerForm.StartIndex:rankerForm.EndIndex]
	context.JSON(http.StatusOK, gin.H{"data": selectedRankings})
}

// Get team ranking by submission id
func (ctrl RankerController) GetTeamRankingsBySubmissionID(context *gin.Context) {

	_, err := fetchAndValidateEvent(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid event name"})
		context.Abort()
		return
	}

	var rankerFormByID forms.RankerFormByID
	if context.ShouldBindJSON(&rankerFormByID) != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
		context.Abort()
		return
	}

	submissionRanking, err := submissionRankingModel.OneBySubmissionID(rankerFormByID.SubmissionID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"Message": "Error fetching submission ranking", "error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": submissionRanking})
}
