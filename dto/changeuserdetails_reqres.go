package dto

type ChangeUserDetailsReqRes struct {
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	Contact        string `json:"contact,omitempty"`
	ProfilePicture []byte `json:"profilePicture,omitempty"`
	ReferralNumber int    `json:"referralNumber,omitempty"`
	AccountNumber  int    `json:"accountNumber,omitempty"`
}
