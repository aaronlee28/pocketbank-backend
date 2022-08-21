package dto

type UserDetailsRes struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Contact        string `json:"contact"`
	ProfilePicture []byte `json:"profilePicture"`
	ReferralNumber int    `json:"referralNumber"`
	AccountNumber  int    `json:"accountNumber"`
}
