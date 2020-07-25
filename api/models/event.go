package models

type Event string

const (
	TestEvent     Event = "testevent"
	HackNRoll2021 Event = "hacknroll2021"
)

var Events = map[string]bool{
	string(TestEvent):     true,
	string(HackNRoll2021): true,
}
