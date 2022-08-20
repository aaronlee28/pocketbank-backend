package dto

type ChangeUserDetailsReqRes struct {
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	Contact        string `json:"contact,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	ReferralNumber int    `json:"referral_number,omitempty"`
	AccountNumber  int    `json:"account_number,omitempty"`
}
