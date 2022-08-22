package models

type Promotion struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Photo []byte `json:"photo"`
}
