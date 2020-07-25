package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/rank-a-thon/rank-a-thon/api/database"
)

// Submission ...
type SubmissionLike struct {
	gorm.Model
	UserID       uint `gorm:"column:user_id;not null" json:"-"`
	SubmissionID uint `gorm:"column:submission_id;not null" json:"-"`
}

// SubmissionModel ...
type SubmissionLikeModel struct{}

// Create ...
func (m SubmissionLikeModel) Create(submissionID uint, userID uint) (likeID uint, err error) {
	_, err = m.One(submissionID, userID)
	if err == nil {
		return 0, errors.New("like already exists")
	}

	submissionLike := SubmissionLike{
		UserID:       userID,
		SubmissionID: submissionID,
	}
	err = database.GetDB().Table("public.submission_likes").Create(&submissionLike).Error
	return submissionLike.ID, err
}

// One ...
func (m SubmissionLikeModel) One(submissionID uint, userID uint) (submissionLike SubmissionLike, err error) {
	err = database.GetDB().Table("public.submissions_likes").
		Where("submission_likes.submissionID = ?  AND submission_likes.user_id = ?", submissionID, userID).
		Take(&submissionLike).Error
	return submissionLike, err
}

// All ...
func (m SubmissionLikeModel) AllForUserID(userID uint) (submissionLikes []SubmissionLike, err error) {
	err = database.GetDB().Table("public.submission_likes").
		Where("submission_likes.user_id = ?", userID).
		Order("submission_likes.id desc").
		Find(&submissionLikes).Error
	return submissionLikes, err
}

// Delete ...
func (m SubmissionLikeModel) Delete(submissionID uint, userID uint) (err error) {
	err = database.GetDB().Table("public.submission_likes").
		Where("submission_likes.submission_id = ?  AND submission_likes.user_id = ?", submissionID, userID).
		Delete(SubmissionLike{}).Error
	return err
}
