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

	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/rank-a-thon/rank-a-thon/api/models"

	"github.com/stretchr/testify/assert"
)


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
