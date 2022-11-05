package dto

type UserDetailsRes struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Contact        string `json:"contact"`
	ProfilePicture string `json:"profilePicture"`
	ReferralNumber int    `json:"referralNumber"`
	SavingsNumber  int    `json:"accountNumber"`
}
