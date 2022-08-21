package models

type User struct {
	Id                  int    `json:"id" gorm:"primarykey"`
	Name                string `json:"name"`
	Email               string `json:"email"`
	Contact             string `json:"contact"`
	Password            string `json:"password,omitempty"`
	Code                int    `json:"code"`
	ReferralNumber      int    `json:"referralNumber"`
	ProfilePicture      []byte `json:"profilePicture"`
	EligibleMerchandise bool   `json:"eligibleMerchandise"`
	Role                string `json:"role"`
}
