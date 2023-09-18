package models

type Transaction struct {
	ID          string  `json:"id"`
	UserID      string  `json:"userId"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	IsFixed     bool    `json:"is_fixed"`
	DayOfMonth  int     `json:"day_of_month"`
	EndDate     string  `json:"endDate"`
	Category    string  `json:"category"`
}
