package model

type Comment struct {
	ID              string
	ProblemID       string
	UserID          string
	Content         string
	ParentCommentID string
	Replies         []Comment
}
