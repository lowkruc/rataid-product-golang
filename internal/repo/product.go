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

	// TODO: replace with database limit and offset
	// for next indicator, increase value param.Limit +1
	// and check if result length data from db is same param.Limit set next is true
	// if data is less then param.Limit + 1, set next is false
	offset := (param.Page - 1) * param.Limit
	if offset >= len(mockProducts) {
		mockProducts = []entity.Product{}
	} else if param.Limit < len(mockProducts) {
		hasNext = true
		startIndex := 0
		limitIndex := param.Limit
		if offset > 0 {
			startIndex = offset
			limitIndex = (param.Limit) + offset
		}

		if limitIndex >= len(mockProducts) {
			hasNext = false
			limitIndex = len(mockProducts)
		}

		mockProducts = mockProducts[startIndex:limitIndex]
	}

	return entity.GetProductResponse{
		Page:     param.Page,
		Next:     hasNext,
		Products: mockProducts,
	}, nil
}
