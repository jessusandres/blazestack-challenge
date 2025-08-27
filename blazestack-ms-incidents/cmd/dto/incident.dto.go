package dto

type ProductToCreate struct {
	Name string `json:"name" binding:"required,min=2,max=255"`
	Sku  string `json:"sku" binding:"required,min=2,max=255"`
}
