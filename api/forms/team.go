package forms

type TeamForm struct {
	TeamName				string	  `form:"team_name" json:"team_name" binding:"required,max=100"`
	IsFreshmanTeam          bool      `form:"is_freshman_team" json:"is_freshman_team"`
	IsPreUniversityTeam     bool      `form:"is_pre_university_team" json:"is_pre_university_team"`
	IsBeginnerTeam          bool      `form:"is_beginner_team" json:"is_beginner_team"`
}