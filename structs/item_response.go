package structs

type ItemResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}
