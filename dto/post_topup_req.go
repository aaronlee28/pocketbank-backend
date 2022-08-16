package dto

type TopupReq struct {
	Amount          int `json:"amount" binding:"required"`
	SourceOfFundsID int `json:"source_of_funds" binding:"required"`
}
