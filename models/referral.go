package models

type Referral struct {
	Id             int `json:"id" gorm:"primarykey"`
	ReferralNumber int `json:"referralNumber"`
	UsedByUserID   int `json:"usedByUserID"`
}
