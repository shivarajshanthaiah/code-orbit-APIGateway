package model

type TestCase struct {
	Input          string `json:"input" bson:"input"`                     // Single input as string
	ExpectedOutput string `json:"expected_output" bson:"expected_output"` // Expected output as string
}

// type TestCasesDocument struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty"`                // MongoDB ObjectID
// 	ProblemID int                `json:"problem_id" bson:"problem_id"` // Problem ID (integer)
// 	TestCases []TestCase         `json:"test_cases" bson:"test_cases"` // Array of test cases
// }
