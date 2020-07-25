package forms

type EvaluationForm struct {
	MainRating             uint `form:"main_rating" json:"main_rating" binding:"required"`
	AnnoyingRating         uint `form:"column:annoying_rating" json:"annoying_rating" binding:"required"`
	EntertainRating        uint `form:"column:entertaining_rating" json:"entertaining_rating" binding:"required"`
	BeautifulRating        uint `form:"column:beautiful_rating" json:"beautiful_rating" binding:"required"`
	SociallyUsefulRating   uint `form:"column:socially_useful_rating" json:"socially_useful_rating" binding:"required"`
	HardwareRating         uint `form:"column:hardware_rating" json:"hardware_rating" binding:"required"`
	AwesomelyUselessRating uint `form:"column:awesomely_useless_rating" json:"awesomely_useless_rating" binding:"required"`
}

type EvaluationFormFloat struct {
	MainRating             float64 `form:"main_rating" json:"main_rating" binding:"required"`
	AnnoyingRating         float64 `form:"column:annoying_rating" json:"annoying_rating" binding:"required"`
	EntertainRating        float64 `form:"column:entertaining_rating" json:"entertaining_rating" binding:"required"`
	BeautifulRating        float64 `form:"column:beautiful_rating" json:"beautiful_rating" binding:"required"`
	SociallyUsefulRating   float64 `form:"column:socially_useful_rating" json:"socially_useful_rating" binding:"required"`
	HardwareRating         float64 `form:"column:hardware_rating" json:"hardware_rating" binding:"required"`
	AwesomelyUselessRating float64 `form:"column:awesomely_useless_rating" json:"awesomely_useless_rating" binding:"required"`
}
