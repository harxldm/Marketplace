package model

import "time"

type Product struct {
	ProductID int       `json:"productID"`
	Name      string    `json:"name"`
	SKU       string    `json:"sku"`
	Amount    int       `json:"amount"`
	Price     float64   `json:"price"`
	SellerID  int       `json:"seller_id"` // El ID del vendedor que crea el producto
	CreatedAt time.Time `json:"created_at"`
}
