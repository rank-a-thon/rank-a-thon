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
		ID     int
		Message string
	}{}

	json.Unmarshal(body, &res)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, 1, res.ID)
}

