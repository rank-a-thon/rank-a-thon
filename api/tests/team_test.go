package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rank-a-thon/rank-a-thon/api/models"
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
	r.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := struct {
		Message string
		ID      int
	}{}

	json.Unmarshal(body, &res)

	teamID = uint(res.ID)

	assert.Equal(t, 1, res.ID)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Team created", res.Message)
}

func TestGetTeam(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/team/testevent", nil)
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
		Data models.Team
	}{}

	json.Unmarshal(body, &res)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Test Team", res.Data.TeamName)
	assert.Equal(t, true, res.Data.IsBeginnerTeam)
	assert.Equal(t, false, res.Data.IsPreUniversityTeam)
	assert.Equal(t, false, res.Data.IsFreshmanTeam)
}

func TestUpdateTeam(t *testing.T) {
	var teamForm forms.TeamForm

	teamForm.TeamName = "Test Team Update"
	teamForm.IsBeginnerTeam = false
	teamForm.IsPreUniversityTeam = true
	teamForm.IsFreshmanTeam = false

	data, _ := json.Marshal(teamForm)

	req, err := http.NewRequest("PUT", "/v1/team/testevent", bytes.NewBufferString(string(data)))
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
	assert.Equal(t, "Team updated", res.Message)
}

const user2Email string = "test2@test.com"
const user2Password string = "123456"


func TestSendTeamInvite(t *testing.T) {
	Register(r, "user2", user2Email, user2Password)
	Login(r, testEmail, testPassword)

	var inviteForm forms.TeamInviteForm

	inviteForm.Event = "testevent"
	inviteForm.Email = user2Email

	data, _ := json.Marshal(inviteForm)

	req, err := http.NewRequest("POST", "/v1/team-invite", bytes.NewBufferString(string(data)))
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
		Message   string
	}{}

	json.Unmarshal(body, &res)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Team invite created", res.Message)
}

func TestGetTeamInvite(t *testing.T) {
	Login(r, user2Email, user2Password)

	req, err := http.NewRequest("GET", "/v1/team-invites", nil)
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
		Data   []models.TeamInvite
	}{}

	json.Unmarshal(body, &res)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, uint(2), res.Data[0].UserID)
	assert.Equal(t, teamID, res.Data[0].TeamID)
}

func TestAcceptTeamInvite(t *testing.T) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/v1/team-invite/accept?teamid=%d", teamID), nil)
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
	assert.Equal(t, "Team invite accepted", res.Message)
}
