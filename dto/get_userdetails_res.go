package dto

type UserDetailsRes struct {
	Id           int     `json:"id"`
	Email        string  `json:"email"`
	WalletID     int     `json:"wallet_id"`
	WalletNumber int     `json:"wallet_number"`
	Balance      float32 `json:"balance"`
}
