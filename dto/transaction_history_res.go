package dto

type TransHistoryRes struct {
	TotalLength  int        `json:"totalLength"`
	Transactions []TransRes `json:"transactions"`
}
