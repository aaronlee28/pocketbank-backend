package dto

type PromotionReq struct {
	Title string `json:"title" binding:"required"`
	Photo string `json:"photo" binding:"required"`
}
