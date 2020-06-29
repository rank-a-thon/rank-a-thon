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

var teamID uint

/**
* TestCreateTeam
* Test team creation
*
* Must return response code 200
 */
func TestCreateTeam(t *testing.T) {
	testRouter := SetupRouter()
	Register(testRouter)
	Login(testRouter)

	var teamForm forms.TeamForm

	teamForm.TeamName = "Test Team"
	teamForm.IsBeginnerTeam = true
	teamForm.IsPreUniversityTeam = false
	teamForm.IsFreshmanTeam = false

	data, _ := json.Marshal(teamForm)

	req, err := http.NewRequest("POST", "/v1/team/testevent", bytes.NewBufferString(string(data)))
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
		Message string
		ID      uint
	}{}

	json.Unmarshal(body, &res)

	teamID = res.ID

	assert.Equal(t, resp.Code, http.StatusOK)
}