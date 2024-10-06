package model

type Subscription struct {
	Plan       string  `json:"plan"`
	Duration   string  `json:"duration"`
	Price      float64 `json:"price"`
	GST        float64 `json:"gst"`
	TotalPrice float64 `json:"total_price"`
}
