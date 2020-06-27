package models

import (
	//"errors"
	"github.com/jinzhu/gorm"
	//"github.com/rank-a-thon/rank-a-thon/api/database"
	//"github.com/rank-a-thon/rank-a-thon/api/forms"
)

// Team
type Team struct {
	gorm.Model
	Users 					[]User  `` //TODO
	IsFreshmanTeam          bool    `gorm:"column:is_freshman_team;default:false" json:"is_freshman_team"`
	IsPreUniversityTeam     bool    `gorm:"column:is_pre_university_team;default:false" json:"is_pre_university_team"`
	IsBeginnerTeam          bool    `gorm:"column:is_beginner_team;default:false" json:"is_beginner_team"`
}

// TeamModel
type TeamModel struct {}

//TODO