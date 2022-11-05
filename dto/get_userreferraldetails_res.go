package dto

type UserReferralDetailsRes struct {
	Count       int   `json:"count"`
	ListOfUsers []int `json:"listOfUsers"`
}
