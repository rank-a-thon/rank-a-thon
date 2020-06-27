package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/rank-a-thon/rank-a-thon/api/database"
)

// TeamInvite ...
type TeamInvite struct {
	gorm.Model
	TeamID      uint      `gorm:"column:team_id;not null" json:"team_id"`
	UserID      uint      `gorm:"column:user_id;not null" json:"user_id"`
}

// TeamInviteModel ...
type TeamInviteModel struct{}

// Create ...
func (m TeamInviteModel) Create(userID uint, teamID uint) (teamInviteID uint, err error) {
	teamInvite := TeamInvite{TeamID: teamID, UserID: userID}
	err = database.GetDB().Table("public.team_invites").Create(&teamInvite).Error
	return teamInvite.ID, err
}

// One ...
func (m TeamInviteModel) One(userID, teamID uint) (teamInvite TeamInvite, err error) {
	err = database.GetDB().Table("public.team_invites").
		Where("team_invites.user_id = ? AND team_invites.team_id = ?", userID, teamID).
		Take(&teamInvite).Error
	return teamInvite, err
}

// All ...
func (m TeamInviteModel) All(userID uint) (teamInvites []TeamInvite, err error) {
	err = database.GetDB().Table("public.team_invites").
		Where("team_invites.user_id = ?", userID).
		Order("team_invites.id desc").
		Find(&teamInvites).Error
	return teamInvites, err
}

// Delete ...
func (m TeamInviteModel) Delete(userID, teamID uint) (err error) {
	teamInvite, err := m.One(userID, teamID)

	if err != nil {
		return errors.New("team invite not found")
	}
	err = database.GetDB().Table("team_invites").Where("id = ?", teamInvite.ID).Delete(TeamInvite{}).Error

	return err
}

