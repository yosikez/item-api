package structs

type ItemRequest struct {
	Name     string `json:"name" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}
