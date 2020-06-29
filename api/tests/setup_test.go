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
	)
	os.Exit(exitVal)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.TestMode)

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)
		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)
		v1.GET("/user/logout", user.Logout)
		v1.GET("/user", user.One)

		/*** START Team ***/
		team := new(controllers.TeamController)
		v1.POST("/team/:event", team.Create)
		v1.GET("/teams", team.All)
		v1.GET("/team/:event", team.One)
		v1.PUT("/team/:event", team.Update)
		v1.DELETE("/team/:event", team.Delete)
		v1.DELETE("/remove-team-member/:event", team.RemoveTeamMember)

		/*** Team Invites ***/
		v1.POST("/team-invite", team.SendInvite)
		v1.GET("/team-invites", team.GetInvites)
		v1.DELETE("/team-invite/accept", team.AcceptInvite)
		v1.DELETE("/team-invite/decline", team.DeclineInvite)

		/*** START AUTH ***/
		auth := new(controllers.AuthController)
		// Refresh the token when needed to generate new access_token and refresh_token for the user
		v1.POST("/token/refresh", auth.Refresh)

		/*** START Submission ***/
		submission := new(controllers.SubmissionController)
		v1.POST("/submission/:event", submission.Create)
		v1.GET("/submissions", submission.AllForUserID)
		v1.GET("/submissions/:event", submission.AllForEvent)
		v1.GET("/submission/:event", submission.One)
		v1.PUT("/submission/:event", submission.Update)
		v1.DELETE("/submission/:event", submission.Delete)

		/*** START Evaluation ***/
		evaluation := new(controllers.EvaluationController)
		v1.GET("/evaluations", evaluation.All)
		v1.GET("/evaluation/:id", evaluation.One)
		v1.PUT("/evaluation/:id", evaluation.Update)
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
