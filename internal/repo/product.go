package repo

import (
	"context"

	"github.com/lowkruc/rataid-product-golang.git/internal/entity"
)

func (db *DBRepo) GetProducts(ctx context.Context, param entity.GetProductRequest) (result entity.GetProductResponse, err error) {
	// TODO: get product to database
	mockProducts := []entity.Product{
		{
			IDProduct:    1,
			Name:         "Product 1",
			Price:        12000,
			Description:  "Product test 1",
			Category:     "Category 1",
			IsOutOfStock: false,
		},
		{
			IDProduct:    2,
			Name:         "Product 2",
			Price:        14000,
			Description:  "Product test 2",
			Category:     "Category 2",
			IsOutOfStock: false,
		},
		{
			IDProduct:    3,
			Name:         "Product 3",
			Price:        16000,
			Description:  "Product test 3",
			Category:     "Category 3",
			IsOutOfStock: true,
		},
	}

	// indicator next page
	hasNext := false

	// mock params implementations
	// becuse length 3 on mock data product, return empty if limit more then 3 and page more then 1
	if param.Limit > 3 && param.Page > 1 {
		mockProducts = []entity.Product{}
	}

	if param.Limit < 3 && param.Page == 1 {
		hasNext = true
		mockProducts = mockProducts[:param.Limit]
	}

	return entity.GetProductResponse{
		Page:     param.Page,
		Next:     hasNext,
		Products: mockProducts,
	}, nil

	return
}
