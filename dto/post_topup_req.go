package dto

type TopupReq struct {
	Amount          float32 `json:"amount" binding:"required"`
	SourceOfFundsID int     `json:"source_of_funds" binding:"required"`
}
