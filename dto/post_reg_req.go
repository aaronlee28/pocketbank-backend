package dto

type RegReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Contact  string `json:"contact"`
	Password string `json:"password,omitempty"`
}
