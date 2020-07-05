package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"github.com/rank-a-thon/rank-a-thon/api/controllers"
	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/models"
	uuid "github.com/twinj/uuid"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// CORSMiddleware ...
// CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// RequestIDMiddleware ...
// Generate a unique ID and attach it to each request for future reference or use
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		newUUID := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", newUUID.String())
		c.Next()
	}
}

var auth = new(controllers.AuthController)

// TokenAuthMiddleware ...
// JWT Authentication middleware attached to each request that needs to be authenticated to validate the access_token in the header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}

func autoMigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.TeamInvite{},
		&models.Submission{},
		&models.Evaluation{},
	).Error
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Start the default gin server
	r := gin.Default()

	// Load the .env file
	envLoadError := godotenv.Load()
	if envLoadError != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	_, fireBaseErr := firebase.NewApp(context.Background(), nil, opt) // app, err :=
	if fireBaseErr != nil {
		fmt.Errorf("error initializing app: %v", fireBaseErr)
		return
	}

	r.Use(CORSMiddleware())
	r.Use(RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Start PostgreSQL database
	// Example: db.GetDB() - More info in the models folder
	database.Init()
	db := database.GetDB()
	autoMigrateDB(db)
	defer db.Close()

	// Start Redis on database 1 - it's used to store the JWT but you can use it for anythig else
	// Example: db.GetRedis().Set(KEY, VALUE, at.Sub(now)).Err()
	database.InitRedis("1")

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
		v1.GET("/evaluations", TokenAuthMiddleware(), evaluation.All)
		v1.GET("/evaluation/:id", TokenAuthMiddleware(), evaluation.One)
		v1.PUT("/evaluation/:id", TokenAuthMiddleware(), evaluation.Update)

		/*** START Ranker ***/
		ranker := new(controllers.RankerController)
		v1.GET("/ranker/start-evaluations", TokenAuthMiddleware(), ranker.CreateEvaluations)
		v1.GET("/ranker/get-team-rankings", TokenAuthMiddleware(), ranker.GetTeamRankings)

	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	fmt.Println("SSL", os.Getenv("SSL"))
	port := os.Getenv("PORT")

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	if os.Getenv("SSL") == "TRUE" {

		SSLKeys := &struct {
			CERT string
			KEY  string
		}{}

		// Generated using sh generate-certificate.sh
		SSLKeys.CERT = "./cert/myCA.cer"
		SSLKeys.KEY = "./cert/myCA.key"

		r.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		r.Run(":" + port)
	}

}
