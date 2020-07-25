package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rank-a-thon/rank-a-thon/api/models"
	"strings"

	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rank-a-thon/rank-a-thon/api/forms"
	"github.com/stretchr/testify/assert"
)

/**
* TestCreateSubmission
* Test submission creation
*
* Must return response code 200
 */
func TestCreateSubmission(t *testing.T) {
	var submissionForm forms.SubmissionForm

	submissionForm.ProjectName = "Testing submission project name"
	submissionForm.Description = "Testing submission description"
	submissionForm.Images = []string{"image1url", "image2url"}

	data, _ := json.Marshal(submissionForm)

	req, err := http.NewRequest("POST", "/v1/submission/testevent", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := struct {
		ID      int
		Message string
	}{}

	json.Unmarshal(body, &res)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, 1, res.ID)
}

func TestCreateSubmissionInvalidAuth(t *testing.T) {
	var submissionForm forms.SubmissionForm

	submissionForm.ProjectName = "Testing submission project name"
	submissionForm.Description = "Testing submission description"
	submissionForm.Images = []string{"image1url", "image2url"}

	data, _ := json.Marshal(submissionForm)

	req, err := http.NewRequest("POST", "/v1/submission/testevent", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", "rubbish"))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := struct {
		ID      int
		Message string
	}{}

	json.Unmarshal(body, &res)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Equal(t, 0, res.ID)
}

func TestGetSubmission(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/submission/testevent", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := struct {
		Data models.Submission
	}{}

	err = json.Unmarshal(body, &res)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, uint(1), res.Data.ID)
	assert.Equal(t, "Testing submission project name", res.Data.ProjectName)
	assert.Equal(t, "Testing submission description", res.Data.Description)
	assert.Equal(t, []string{"image1url", "image2url"}, strings.Split(res.Data.Images, ","))
}

func TestUpdateSubmission(t *testing.T) {
	var submissionForm forms.SubmissionForm

	submissionForm.ProjectName = "New submission project name"
	submissionForm.Description = "New submission description"
	submissionForm.Images = []string{"newimage1url"}

	data, _ := json.Marshal(submissionForm)

	req, err := http.NewRequest("PUT", "/v1/submission/testevent", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := struct {
		Message string
	}{}

	json.Unmarshal(body, &res)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Submission updated", res.Message)
}

func TestDeleteSubmission(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/submission/testevent", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := struct {
		Message string
	}{}

	json.Unmarshal(body, &res)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Submission deleted", res.Message)
}
