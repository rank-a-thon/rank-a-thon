package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/rank-a-thon/rank-a-thon/api/database"
	"github.com/rank-a-thon/rank-a-thon/api/forms"
)

// Submission ...
type Submission struct {
	gorm.Model
	UserID    uint   `gorm:"column:user_id;not null" json:"-"`
	Title     string `gorm:"column:title" json:"title"`
	Content   string `gorm:"column:content" json:"content"`
	User      User   `gorm:"column:user;foreignkey:UserID" json:"user"`
}

// SubmissionModel ...
type SubmissionModel struct{}

// Create ...
func (m SubmissionModel) Create(userID uint, form forms.SubmissionForm) (submissionID uint, err error) {
	submission := Submission{UserID: userID, Title: form.Title, Content: form.Content}
	err = database.GetDB().Table("public.submissions").Create(&submission).Error
	return submission.ID, err
}

// One ...
func (m SubmissionModel) One(userID, id uint) (submission Submission, err error) {
	err = database.GetDB().Preload("User").Table("public.submissions").
		Where("submissions.user_id = ? AND submissions.id = ?", userID, id).
		//Select("submissions.id, submissions.title, submissions.content, submissions.updated_at, submissions.created_at").
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
		Updates(map[string]interface{}{"title": form.Title, "content": form.Content}).Error
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
