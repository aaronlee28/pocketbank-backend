package dto

type PromotionReq struct {
	Title string `json:"title" binding:"required"`
	Photo []byte `json:"photo" binding:"required"`
}
