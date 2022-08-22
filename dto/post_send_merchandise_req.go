package dto

type MerchandiseStatus struct {
	UserID      int    `json:"userID"`
	MerchToSend string `json:"merchToSend"`
	Status      string `json:"status"`
}
