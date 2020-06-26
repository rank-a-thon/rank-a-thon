package forms

type EvaluationForm struct {
	MainRanking		        uint    `form:"main_ranking" json:"main_ranking" binding:"required"`
	AnnoyingRanking         uint    `form:"column:annoying_ranking" json:"annoying_ranking" binding:"required"`
	EntertainRanking        uint	`form:"column:entertaining_ranking" json:"entertaining_ranking" binding:"required"`
	BeautifulRanking        uint	`form:"column:beautiful_ranking" json:"beautiful_ranking" binding:"required"`
	SociallyUsefulRanking   uint    `form:"column:socially_useful_ranking" json:"socially_useful_ranking" binding:"required"`
	HardwareRanking    	    uint    `form:"column:hardware_ranking" json:"hardware_ranking" binding:"required"`
	AwesomelyUselessRanking uint	`form:"column:awesomely_useless_ranking" json:"awesomely_useless_ranking" binding:"required"`
}
