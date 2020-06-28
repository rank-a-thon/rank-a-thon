package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/forms"
)

// Team
type Team struct {
	gorm.Model
	TeamName				string	  `gorm:"column:team_name;not_null" json:"team_name"`
	Event					Event	  `gorm:"column:event;not_null" json:"event"`
	UserIDs 				string    `gorm:"column:user_ids;not_null" json:"user_ids"`
	IsFreshmanTeam          bool      `gorm:"column:is_freshman_team;default:false" json:"is_freshman_team"`
	IsPreUniversityTeam     bool      `gorm:"column:is_pre_university_team;default:false" json:"is_pre_university_team"`
	IsBeginnerTeam          bool      `gorm:"column:is_beginner_team;default:false" json:"is_beginner_team"`
}

// TeamModel
type TeamModel struct {}

// Create ...
func (m TeamModel) Create(userID uint, form forms.TeamForm, event Event) (TeamID uint, err error) {
	teamID, err := getTeamIDForUser(userID, event)
	if err != nil {
		return 0, err
	} else if teamID != 0 {
		return teamID, errors.New(fmt.Sprintf("error creating team: user %d already belongs to team %d", userID, teamID))
	}
	team := Team{
		TeamName:            form.TeamName,
		Event:				 event,
		UserIDs:             UintSliceToJsonString([]uint{userID}),
		IsFreshmanTeam:      form.IsFreshmanTeam,
		IsPreUniversityTeam: form.IsPreUniversityTeam,
		IsBeginnerTeam:      form.IsBeginnerTeam,
	}
	err = database.GetDB().Table("public.teams").Create(&team).Error
	err = userModel.UpdateTeamForUser(userID, team.ID, event)
	return team.ID, err
}

// One ...
func (m TeamModel) One(teamID uint) (team Team, err error) {
	err = database.GetDB().Table("public.teams").
		Where("teams.id = ?", teamID).
		Take(&team).Error
	return team, err
}

// All ...
func (m TeamModel) All(userID uint) (teams []Team, err error) {
	user, err := userModel.One(userID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user %d not found", userID))
	}
	for _, teamID := range JsonStringToStringUintMap(user.TeamIDForEvent) {
		team, _ := m.One(teamID)
		teams = append(teams, team)
	}
	return teams, err
}

// Update ...
func (m TeamModel) Update(teamID uint, form forms.TeamForm) (err error) {
	_, err = m.One(teamID)

	if err != nil {
		return errors.New(fmt.Sprintf("team %d not found", teamID))
	}
	err = database.GetDB().Table("public.teams").Model(&Team{}).
		Where("id = ?", teamID).
		Updates(map[string]interface{}{
			"team_name": form.TeamName,
			"is_freshman_team": form.IsFreshmanTeam,
			"is_pre_university_team": form.IsPreUniversityTeam,
			"is_beginner_team": form.IsBeginnerTeam,
		}).Error
	return err
}

func (m TeamModel) AddTeamMember(userID uint, teamID uint) (err error) {
	team, err := m.One(teamID)
	if err != nil {
		return errors.New(fmt.Sprintf("team %d not found", teamID))
	}

	user, err := userModel.One(userID)
	if err != nil {
		return errors.New(fmt.Sprintf("user %d not found", userID))
	}
	if JsonStringToStringUintMap(user.TeamIDForEvent)[string(team.Event)] != 0 {
		return errors.New(fmt.Sprintf("user %d is already part of team %d for event %s", userID, teamID, team.Event))
	}

	teamUserIDs := JsonStringToUintSlice(team.UserIDs)
	for _, id := range teamUserIDs {
		if id == userID {
			return errors.New(fmt.Sprintf("user %d is already part of team %d for event %s", userID, teamID, team.Event))
		}
	}
	err = database.GetDB().Table("public.teams").Model(&Team{}).
		Where("id = ?", teamID).
		Updates(map[string]interface{}{
			"user_ids": UintSliceToJsonString(append(teamUserIDs, userID)),
		}).Error
	err = userModel.UpdateTeamForUser(userID, teamID, team.Event)
	return err
}

func (m TeamModel) RemoveTeamMember(userID uint, teamID uint) (err error) {
	team, err := m.One(teamID)

	if err != nil {
		return errors.New(fmt.Sprintf("team %d not found", teamID))
	}

	indexToRemove := -1
	teamUserIDs := JsonStringToUintSlice(team.UserIDs)
	for i, v := range teamUserIDs {
		if v == userID {
			indexToRemove = i
		}
	}
	if indexToRemove == -1 {
		return errors.New(fmt.Sprintf("user %d does not belong in team %d", userID, teamID))
	}

	err = database.GetDB().Table("public.teams").Model(&Team{}).
		Where("id = ?", teamID).
		Updates(map[string]interface{}{
			"user_ids": append(teamUserIDs[:indexToRemove], teamUserIDs[indexToRemove+1:]...),
		}).Error
	err = userModel.UpdateTeamForUser(userID, 0, team.Event)
	return err
}

// Delete ...
func (m TeamModel) Delete(teamID uint) (err error) {
	team, err := m.One(teamID)
	if err != nil {
		return errors.New(fmt.Sprintf("team %d not found", teamID))
	}
	teamUserIDs := JsonStringToUintSlice(team.UserIDs)
	for _, id := range teamUserIDs {
		err = userModel.UpdateTeamForUser(id, 0, team.Event)
	}
	err = database.GetDB().Table("public.teams").Where("id = ?", teamID).Delete(Team{}).Error

	return err
}

var userModel = new(UserModel)

func getTeamIDForUser(userID uint, event Event) (uint, error) {
	user, err := userModel.One(userID)
	if err != nil {
		return 0, err
	}
	teamIDForEvent := JsonStringToStringUintMap(user.TeamIDForEvent)
	return teamIDForEvent[string(event)], nil
}