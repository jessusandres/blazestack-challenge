package types

type Product struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Sku     string `json:"sku"`
	Visible bool   `json:"visibleIntoWeb"`
}

type ProductCreatedResponse struct {
	Product Product `json:"product"`
}
