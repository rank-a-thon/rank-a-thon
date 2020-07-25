package models

type UserType int

const (
	Participant UserType = iota // 0
	Sponsor                     // 1
	Judge                       // 2
	Admin                       // 3
)
