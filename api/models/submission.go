package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/forms"
)

// Submission ...
type Submission struct {
	gorm.Model
	TeamID      uint   `gorm:"column:team_id;not null" json:"-"`
	ProjectName string `gorm:"column:project_name" json:"project_name"`
	Description string `gorm:"column:description" json:"description"`
	Images      string `gorm:"column:images" json:"images"` // comma separated list of image ids
	Team        Team   `gorm:"column:team" json:"team"`
}

// SubmissionModel ...
type SubmissionModel struct{}

// Create ...
func (m SubmissionModel) Create(teamID uint, form forms.SubmissionForm) (submissionID uint, err error) {
	submission, err := m.OneByTeamID(teamID)
	if err == nil {
		return submission.ID, errors.New("team submission already exists")
	}

	submission = Submission{
		TeamID:      teamID,
		ProjectName: form.ProjectName,
		Description: form.Description,
		Images:      strings.Join(form.Images, ","),
	}
	err = database.GetDB().Table("public.submissions").Create(&submission).Error
	return submission.ID, err
}

// One ...
func (m SubmissionModel) OneByTeamID(teamID uint) (submission Submission, err error) {
	err = database.GetDB().Preload("Team").Table("public.submissions").
		Where("submissions.team_id = ?", teamID).
		Joins("left join public.teams on submissions.team_id = teams.id").
		Take(&submission).Error
	return submission, err
}

// One ...
func (m SubmissionModel) One(submissionID uint) (submission Submission, err error) {
	err = database.GetDB().Preload("Team").Table("public.submissions").
		Where("submissions.id = ?", submissionID).
		Joins("left join public.teams on submissions.team_id = teams.id").
		Take(&submission).Error
	return submission, err
}

// All ...
func (m SubmissionModel) AllForUserID(userID uint) (submissions []Submission, err error) {
	user, err := userModel.One(userID)
	if err != nil {
		return nil, err
	}
	teamIDForEvent := JsonStringToStringUintMap(user.TeamIDForEvent)
	for _, teamID := range teamIDForEvent {
		submission, err := m.OneByTeamID(teamID)
		if err != nil {
			return submissions, err
		}
		submissions = append(submissions, submission)
	}
	return submissions, err
}

func (m SubmissionModel) AllForEvent(event Event) (submissions []Submission, err error) {
	err = database.GetDB().Preload("Team").Table("public.submissions").
		Joins("left join public.teams on submissions.team_id = teams.id").
		Where("teams.event = ?", event).
		Order("submissions.id desc").
		Find(&submissions).Error
	return submissions, err
}

// Update ...
func (m SubmissionModel) Update(teamID uint, form forms.SubmissionForm) (err error) {
	_, err = m.OneByTeamID(teamID)

	if err != nil {
		return errors.New("submission not found")
	}
	err = database.GetDB().Table("public.submissions").Model(&Submission{}).
		Where("team_id = ?", teamID).
		Updates(map[string]interface{}{
			"project_name": form.ProjectName,
			"description": form.Description,
			"images": strings.Join(form.Images, ","),
		}).Error
	return err
}

// Delete ...
func (m SubmissionModel) Delete(teamID uint) (err error) {
	_, err = m.OneByTeamID(teamID)

	if err != nil {
		return errors.New("submission not found")
	}
	err = database.GetDB().Table("public.submissions").Where("team_id = ?", teamID).Delete(Submission{}).Error

	return err
}