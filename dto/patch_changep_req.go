package dto

type ChangePReq struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
	Code        int    `json:"code"`
}
