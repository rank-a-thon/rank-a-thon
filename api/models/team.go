package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/forms"
)

// Team
type Team struct {
	gorm.Model
	TeamName				string	  `gorm:"column:team_name" json:"team_name"`
	UserIDs 				[]uint    `gorm:"column:user_ids" json:"user_ids"`
	IsFreshmanTeam          bool      `gorm:"column:is_freshman_team;default:false" json:"is_freshman_team"`
	IsPreUniversityTeam     bool      `gorm:"column:is_pre_university_team;default:false" json:"is_pre_university_team"`
	IsBeginnerTeam          bool      `gorm:"column:is_beginner_team;default:false" json:"is_beginner_team"`
}

// TeamModel
type TeamModel struct {}

// Create ...
func (m TeamModel) Create(userID uint, form forms.TeamForm) (TeamID uint, err error) {
	teamID, err := getTeamIDForUser(userID)
	if err != nil {
		return 0, err
	} else if teamID != 0 {
		return teamID, errors.New(fmt.Sprintf("error creating team: user %d already belongs to team %d", userID, teamID))
	}
	team := Team{
		TeamName: form.TeamName,
		UserIDs: []uint{userID},
		IsFreshmanTeam: form.IsFreshmanTeam,
		IsPreUniversityTeam: form.IsPreUniversityTeam,
		IsBeginnerTeam: form.IsBeginnerTeam,
	}
	err = database.GetDB().Table("public.teams").Create(&team).Error
	return team.ID, err
}

// One ...
func (m TeamModel) One(teamID uint) (team Team, err error) {
	err = database.GetDB().Table("public.teams").
		Where("teams.id = ?", teamID).
		Take(&team).Error
	return team, err
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
	err = database.GetDB().Table("public.teams").Model(&Team{}).
		Where("id = ?", teamID).
		Updates(map[string]interface{}{
			"user_ids": append(team.UserIDs, userID),
		}).Error
	return err
}

func (m TeamModel) RemoveTeamMember(userID uint, teamID uint) (err error) {
	team, err := m.One(teamID)

	if err != nil {
		return errors.New(fmt.Sprintf("team %d not found", teamID))
	}

	indexToRemove := -1
	for i, v := range team.UserIDs {
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
			"user_ids": append(team.UserIDs[:indexToRemove], team.UserIDs[indexToRemove+1:]...),
		}).Error
	return err
}

// Delete ...
func (m TeamModel) Delete(teamID uint) (err error) {
	_, err = m.One(teamID)

	if err != nil {
		return errors.New(fmt.Sprintf("team %d not found", teamID))
	}
	err = database.GetDB().Table("public.teams").Where("id = ?", teamID).Delete(Team{}).Error

	return err
}

var userModel = new(UserModel)

func getTeamIDForUser(userID uint) (uint, error) {
	user, err := userModel.One(userID)
	if err != nil {
		return 0, err
	}
	return user.TeamID, nil
}