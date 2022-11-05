package dto

type RegReq struct {
	Name           string `json:"name" form:"name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Contact        string `json:"contact" binding:"required"`
	Password       string `json:"password" binding:"required"`
	ReferralNumber int    `json:"referralNumber,omitempty"`
	Photo          []byte `json:"photo,omitempty"`
}
