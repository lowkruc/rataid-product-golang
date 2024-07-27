package entity

type Product struct {
	IDProduct    int64   `json:"id_product"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	IsOutOfStock bool    `json:"is_out_of_stock"`
}

type GetProductRequest struct {
	Limit int
	Page  int
}

type GetProductResponse struct {
	Page     int
	Next     bool
	Products []Product
}
