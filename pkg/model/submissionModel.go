package model

type Submission struct {
    ProblemID int    `json:"problem_id"`
    Language  string `json:"language"`
    Code      string `json:"code"`
}
