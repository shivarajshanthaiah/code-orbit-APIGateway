package model

type Problem struct {
	ID          uint
	Title       string
	Description string
	Difficulty  string
	Tags        string
	IsPremium   bool
}
