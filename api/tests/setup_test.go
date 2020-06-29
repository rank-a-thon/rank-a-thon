package tests

import (
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rank-a-thon/rank-a-thon/api/controllers"
	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/models"
)

var testEmail = "test-rankathon@test.com"
var testPassword = "123456"

var accessToken string
var refreshToken string

func main() {
	r := SetupRouter()
	r.Run()
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

		/*** START AUTH ***/
		auth := new(controllers.AuthController)
		v1.POST("/token/refresh", auth.Refresh)

		/*** START submission ***/
		submission := new(controllers.SubmissionController)
		v1.POST("/submission/:event", submission.Create)
		v1.GET("/submissions", submission.AllForUserID)
		v1.GET("/submissions/:event", submission.AllForEvent)
		v1.GET("/submission/:event", submission.One)
		v1.PUT("/submission/:event", submission.Update)
		v1.DELETE("/submission/:event", submission.Delete)
	}
	return r
}

/**
* TestIntDB
* It tests the connection to the database and init the db for this test
*
* Must pass
 */
func TestIntDB(t *testing.T) {

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
}