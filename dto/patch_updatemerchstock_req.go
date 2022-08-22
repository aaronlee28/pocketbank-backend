package dto

type UpdateMerchStocksReq struct {
	Name  string `json:"name" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
}
