package dto

type UserDetailsRes struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Contact        string `json:"contact"`
	ProfilePicture string `json:"profile_picture"`
	ReferralNumber int    `json:"referral_number"`
	AccountNumber  int    `json:"account_number"`
}
