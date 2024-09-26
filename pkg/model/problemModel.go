package model

type Problem struct {
	ID          uint
	Title       string
	Description string
	Difficulty  string
	Type        string
	IsPremium   bool
}
