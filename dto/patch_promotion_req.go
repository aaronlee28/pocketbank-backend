package dto

type PatchPromotionReq struct {
	Title string `json:"title,omitempty"`
	Photo []byte `json:"photo,omitempty"`
}
