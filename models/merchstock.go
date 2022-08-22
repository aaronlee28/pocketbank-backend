package models

type Merchstock struct {
	Id         int    `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	StockCount int    `json:"stockCount"`
}
