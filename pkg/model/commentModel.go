package model

type Comment struct {
	ID              string `json:"id"`
	ProblemID       int    `json:"problem_id"`
	UserID          string
	Content         string `json:"content"`
	ParentCommentID string `json:"parent_comment_id"`
	Replies         []Comment
}
