// +build all

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rank-a-thon/rank-a-thon/api/controllers"
	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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

		v1.POST("/submission", submission.Create)
		v1.GET("/submissions", submission.All)
		v1.GET("/submission/:id", submission.One)
		v1.PUT("/submission/:id", submission.Update)
		v1.DELETE("/submission/:id", submission.Delete)
	}

	return r
}

func main() {
	r := SetupRouter()
	r.Run()
}

var loginCookie string

var testEmail = "test-gin-boilerplate@test.com"
var testPassword = "123456"

var accessToken string
var refreshToken string

var submissionID int

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
	db.AutoMigrate(&models.User{}, &models.Submission{})
}

/**
* TestRegister
* Test user registration
*
* Must return response code 200
 */
func TestRegister(t *testing.T) {
	testRouter := SetupRouter()

	var registerForm forms.RegisterForm

	registerForm.Name = "testing"
	registerForm.Email = testEmail
	registerForm.Password = testPassword

	data, _ := json.Marshal(registerForm)

	req, err := http.NewRequest("POST", "/v1/user/register", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
}

/**
* TestRegisterInvalidEmail
* Test user registration with invalid email
*
* Must return response code 406
 */
func TestRegisterInvalidEmail(t *testing.T) {
	testRouter := SetupRouter()

	var registerForm forms.RegisterForm

	registerForm.Name = "testing"
	registerForm.Email = "invalid@email"
	registerForm.Password = testPassword

	data, _ := json.Marshal(registerForm)

	req, err := http.NewRequest("POST", "/v1/user/register", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusNotAcceptable)
}

/**
* TestLogin
* Test user login
* and get the access_token and refresh_token stored
*
* Must return response code 200
 */
func TestLogin(t *testing.T) {
	testRouter := SetupRouter()

	var loginForm forms.LoginForm

	loginForm.Email = testEmail
	loginForm.Password = testPassword

	data, _ := json.Marshal(loginForm)

	req, err := http.NewRequest("POST", "/v1/user/login", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res = &struct {
		Message string `json:"message"`
		User    struct {
			CreatedAt int64  `json:"created_at"`
			Email     string `json:"email"`
			ID        int64  `json:"id"`
			Name      string `json:"name"`
			UpdatedAt int64  `json:"updated_at"`
		} `json:"user"`
		Token struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"token"`
	}{}

	json.Unmarshal(body, &res)

	accessToken = res.Token.AccessToken
	refreshToken = res.Token.RefreshToken

	assert.Equal(t, resp.Code, http.StatusOK)
}

/**
* TestInvalidLogin
* Test invalid login
*
* Must return response code 406
 */
func TestInvalidLogin(t *testing.T) {
	testRouter := SetupRouter()

	var loginForm forms.LoginForm

	loginForm.Email = "wrong@email.com"
	loginForm.Password = testPassword

	data, _ := json.Marshal(loginForm)

	req, err := http.NewRequest("POST", "/v1/user/login", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusNotAcceptable)
}

/**
* TestCreateSubmission
* Test submission creation
*
* Must return response code 200
 */
func TestCreateSubmission(t *testing.T) {
	testRouter := SetupRouter()

	var submissionForm forms.SubmissionForm

	submissionForm.ProjectName = "Testing submission project name"
	submissionForm.Description = "Testing submission description"

	data, _ := json.Marshal(submissionForm)

	req, err := http.NewRequest("POST", "/v1/submission", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := struct {
		Status int
		ID     int
	}{}

	json.Unmarshal(body, &res)

	submissionID = res.ID

	assert.Equal(t, resp.Code, http.StatusOK)
}

/**
* TestCreateInvalidSubmission
* Test submission invalid creation
*
* Must return response code 406
 */
func TestCreateInvalidSubmission(t *testing.T) {
	testRouter := SetupRouter()

	var submissionForm forms.SubmissionForm

	submissionForm.ProjectName = "Testing submission project name"

	data, _ := json.Marshal(submissionForm)

	req, err := http.NewRequest("POST", "/v1/submission", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusNotAcceptable)
}

/**
* TestGetSubmission
* Test getting one submission
*
* Must return response code 200
 */
func TestGetSubmission(t *testing.T) {
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", fmt.Sprintf("/v1/submission/%d", submissionID), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
}

/**
* TestGetInvalidSubmission
* Test getting invalid submission
*
* Must return response code 404
 */
func TestGetInvalidSubmission(t *testing.T) {
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/v1/submission/invalid", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusNotFound)
}

/**
* TestGetSubmissionNotLoggedin
* Test getting the submission with logged out user
*
* Must return response code 401
 */
func TestGetSubmissionNotLoggedin(t *testing.T) {
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", fmt.Sprintf("/v1/submission/%d", submissionID), nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusUnauthorized)
}

/**
* TestGetSubmissionUnauthorized
* Test getting the submission with unauthorized user (wrong or expired access_token)
*
* Must return response code 401
 */
func TestGetSubmissionUnauthorized(t *testing.T) {
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", fmt.Sprintf("/v1/submission/%d", submissionID), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", "abc123"))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusUnauthorized)
}

/**
* TestUpdateSubmission
* Test updating an submission
*
* Must return response code 200
 */
func TestUpdateSubmission(t *testing.T) {
	testRouter := SetupRouter()

	var submissionForm forms.SubmissionForm

	submissionForm.ProjectName = "Testing new submission project name"
	submissionForm.Description = "Testing new submission description"

	data, _ := json.Marshal(submissionForm)

	url := fmt.Sprintf("/v1/submission/%d", submissionID)

	req, err := http.NewRequest("PUT", url, bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
}

/**
* TestDeleteSubmission
* Test deleting an submission
*
* Must return response code 200
 */
func TestDeleteSubmission(t *testing.T) {
	testRouter := SetupRouter()

	url := fmt.Sprintf("/v1/submission/%d", submissionID)

	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
}

/**
* TestRefreshToken
* Test refreshing the token with valid refresh_token
*
* Must return response code 200
 */
func TestRefreshToken(t *testing.T) {
	testRouter := SetupRouter()

	var tokenForm forms.Token

	tokenForm.RefreshToken = refreshToken

	data, _ := json.Marshal(tokenForm)

	req, err := http.NewRequest("POST", "/v1/token/refresh", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
}

/**
* TestInvalidRefreshToken
* Test refreshing the token with invalid refresh_token
*
* Must return response code 401
 */
func TestInvalidRefreshToken(t *testing.T) {
	testRouter := SetupRouter()

	var tokenForm forms.Token

	//Since we didn't update it in the test before - this will not be valid anymore
	tokenForm.RefreshToken = refreshToken

	data, _ := json.Marshal(tokenForm)

	req, err := http.NewRequest("POST", "/v1/token/refresh", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusUnauthorized)
}

/**
* TestUserSignout
* Test logout a user
*
* Must return response code 200
 */
func TestUserLogout(t *testing.T) {
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/v1/user/logout", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
}

/**
* TestCleanUp
* Deletes the created user with it's submissions
*
* Must pass
 */
func TestCleanUp(t *testing.T) {
	var err error
	err = database.GetDB().Table("public.users").Where("email = ?", testEmail).Delete(&models.User{}).Error
	//_, err = database.GetDB().Exec("DELETE FROM public.user WHERE email=$1", testEmail)
	if err != nil {
		t.Error(err)
	}
}
