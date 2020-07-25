package tests

import (
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rank-a-thon/rank-a-thon/api/controllers"
	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/models"
)

var testEmail = "test-rankathon@test.com"
var testPassword = "123456"

var accessToken string
var refreshToken string
var r *gin.Engine
var auth = new(controllers.AuthController)

func TestMain(m *testing.M) {
	InitDbAndAutoMigrate()
	exitVal := m.Run()
	db := database.GetDB()
	db.DropTableIfExists(
		&models.User{},
		&models.Team{},
		&models.TeamInvite{},
		&models.Submission{},
		&models.Evaluation{},
		&models.SubmissionRanking{},
	)
	os.Exit(exitVal)
}

// TokenAuthMiddleware ...
// JWT Authentication middleware attached to each request that needs to be authenticated to validate the access_token in the header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r = gin.Default()
	gin.SetMode(gin.TestMode)

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)
		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)
		v1.GET("/user/logout", user.Logout)
		v1.GET("/user", user.GetByUserID)

		/*** START Team ***/
		team := new(controllers.TeamController)
		v1.POST("/team/:event", TokenAuthMiddleware(), team.Create)
		v1.GET("/teams", TokenAuthMiddleware(), team.All)
		v1.GET("/team/:event", TokenAuthMiddleware(), team.One)
		v1.PUT("/team/:event", TokenAuthMiddleware(), team.Update)
		v1.DELETE("/team/:event", TokenAuthMiddleware(), team.Delete)
		v1.DELETE("/remove-team-member/:event", TokenAuthMiddleware(), team.RemoveTeamMember)

		/*** Team Invites ***/
		v1.POST("/team-invite", TokenAuthMiddleware(), team.SendInvite)
		v1.GET("/team-invites", TokenAuthMiddleware(), team.GetInvites)
		v1.DELETE("/team-invite/accept", TokenAuthMiddleware(), team.AcceptInvite)
		v1.DELETE("/team-invite/decline", TokenAuthMiddleware(), team.DeclineInvite)

		/*** START AUTH ***/
		auth := new(controllers.AuthController)
		// Refresh the token when needed to generate new access_token and refresh_token for the user
		v1.POST("/token/refresh", auth.Refresh)

		/*** START Submission ***/
		submission := new(controllers.SubmissionController)
		v1.POST("/submission/:event", TokenAuthMiddleware(), submission.Create)
		v1.GET("/submissions", TokenAuthMiddleware(), submission.AllForUserID)
		v1.GET("/submissions/:event", TokenAuthMiddleware(), submission.AllForEvent)
		v1.GET("/submission/:event", TokenAuthMiddleware(), submission.One)
		v1.PUT("/submission/:event", TokenAuthMiddleware(), submission.Update)
		v1.DELETE("/submission/:event", TokenAuthMiddleware(), submission.Delete)

		/*** START Evaluation ***/
		evaluation := new(controllers.EvaluationController)
		v1.GET("/evaluations", TokenAuthMiddleware(), evaluation.AllForJudge)
		v1.GET("/evaluation/:id", TokenAuthMiddleware(), evaluation.One)
		v1.PUT("/evaluation/:id", TokenAuthMiddleware(), evaluation.Update)

		/*** START Ranker ***/
		ranker := new(controllers.RankerController)
		v1.PUT("/ranker/start-evaluations/:event", TokenAuthMiddleware(), ranker.CreateEvaluations)
		v1.PUT("/ranker/calculate-team-rankings/:event", TokenAuthMiddleware(), ranker.CalculateTeamRankings)
		v1.GET("/ranker/team-rankings-by-range/:event", TokenAuthMiddleware(), ranker.GetTeamRankingsByRange)
		v1.GET("/ranker/team-rankings-by-submission-id/:event", TokenAuthMiddleware(), ranker.GetTeamRankingsBySubmissionID)
	}
	return r
}

func InitDbAndAutoMigrate() *gorm.DB {
	//Load the .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

	database.InitForTest()
	database.InitRedis("1")
	db := database.GetDB()
	db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.TeamInvite{},
		&models.Submission{},
		&models.Evaluation{},
	)
	return db
}
