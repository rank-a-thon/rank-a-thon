package forms

type RankerForm struct {
	Category     string    `form:"category" json:"category" binding:"required"`
	StartIndex   uint      `form:"start_index" json:"start_index"`
	EndIndex     uint      `form:"end_index" json:"end_index" binding:"required"`
}

type RankerFormByID struct {
	Category       string    `form:"category" json:"category" binding:"required"`
	SubmissionID   uint      `form:"submission_id" json:"submission_id" binding:"required"`
}