package dto

type ChangePReq struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
	Code        int    `json:"code"`
}
