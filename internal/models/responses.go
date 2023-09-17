package models

type UserTransactionsResponse struct {
	UserID   string         `json:"userID"`
	Incomes  []*Transaction `json:"incomes"`
	Expenses []*Transaction `json:"expenses"`
}
