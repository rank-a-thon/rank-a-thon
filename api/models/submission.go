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
	//TODO change UserID and User to TeamID and Team
	UserID      uint      `gorm:"column:user_id;not null" json:"-"`
	ProjectName string    `gorm:"column:project_name" json:"project_name"`
	Description string    `gorm:"column:description" json:"description"`
	Images      string    `gorm:"column:images" json:"images"` // comma separated list of image ids
	User        User      `gorm:"column:user;foreignkey:UserID" json:"user"`
}

// SubmissionModel ...
type SubmissionModel struct{}

// Create ...
func (m SubmissionModel) Create(userID uint, form forms.SubmissionForm) (submissionID uint, err error) {

	submission := Submission{
		UserID: userID,
		ProjectName: form.ProjectName,
		Description: form.Description,
		Images: strings.Join(form.Images, ","),
	}
	err = database.GetDB().Table("public.submissions").Create(&submission).Error
	return submission.ID, err
}

// One ...
func (m SubmissionModel) One(userID, id uint) (submission Submission, err error) {
	err = database.GetDB().Preload("User").Table("public.submissions").
		Where("submissions.user_id = ? AND submissions.id = ?", userID, id).
		Joins("left join public.users on submissions.user_id = users.id").
		Take(&submission).Error
	return submission, err
}

// All ...
func (m SubmissionModel) All(userID uint) (submissions []Submission, err error) {
	err = database.GetDB().Preload("User").Table("public.submissions").
		Where("submissions.user_id = ?", userID).
		Joins("left join public.users on submissions.user_id = users.id").
		Order("submissions.id desc").
		Find(&submissions).Error
	return submissions, err
}

// Update ...
func (m SubmissionModel) Update(userID uint, id uint, form forms.SubmissionForm) (err error) {
	_, err = m.One(userID, id)

	if err != nil {
		return errors.New("submission not found")
	}
	err = database.GetDB().Table("public.submissions").Model(&Submission{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"project_name": form.ProjectName,
			"description": form.Description,
			"images": strings.Join(form.Images, ","),
		}).Error
	return err
}

// Delete ...
func (m SubmissionModel) Delete(userID, id uint) (err error) {
	_, err = m.One(userID, id)

	if err != nil {
		return errors.New("submission not found")
	}
	err = database.GetDB().Table("public.submissions").Where("id = ?", id).Delete(Submission{}).Error

	return err
}