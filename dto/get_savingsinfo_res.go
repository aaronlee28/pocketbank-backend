package dto

type SavingsRes struct {
	UserID   int     `json:"user_id"`
	Balance  float32 `json:"balance"`
	Interest float32 `json:"interest"`
	Tax      float32 `json:"tax"`
}
