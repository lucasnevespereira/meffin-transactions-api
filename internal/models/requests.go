package models

type CreateTransactionRequest struct {
	UserID      string  `json:"userId"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	IsFixed     bool    `json:"is_fixed"`
	DayOfMonth  int     `json:"day_of_month"`
	EndDate     string  `json:"endDate"`
	CategoryID  string  `json:"category"`
}

type CreateCategoryRequest struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Color  string `json:"color"`
}
