package models

type SaleProduct struct {
	Id                string  `json:"id"`
	ProductId         string  `json:"product_id"`
	ProductName       string  `json:"product_name"`
	ProductPrice      float64 `jsos:"produc_price"`
	Discount          float64 `json:"discount"`
	DiscountType      string  `json:"discount_type"`
	PriceWithDiscount float64 `json:"price_with_discount"`
	DiscountPrice     float64 `json:"discount_price"`
	Count             float64 `json:"count"`
	TotalPrice        float64 `json:"total_price"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}

type CreateSaleProduct struct {
	ProductId    string  `json:"product_id"`
	Discount     float64 `json:"discount"`
	DiscountType string  `json:"discount_type"`
	Count        float64 `json:"count"`
}

type SaleProductPrimaryKey struct {
	Id string `json:"id"`
}
