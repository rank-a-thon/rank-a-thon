package forms

type TeamInviteForm struct {
	Event string `form:"event" json:"event" binding:"required,max=100"`
	Email string `form:"email" json:"email" binding:"required,email"`
}
