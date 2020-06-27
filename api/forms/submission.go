package forms

type SubmissionForm struct {
	ProjectName string `form:"project_name" json:"project_name" binding:"required,max=100"`
	Description string `form:"description" json:"description" binding:"required,max=3000"`
}
