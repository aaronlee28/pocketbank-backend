package models

type Referral struct {
	Id             int `json:"id" gorm:"primarykey"`
	ReferralNumber int `json:"referral_number"`
	UsedCount      int `json:"used_count"`
}

