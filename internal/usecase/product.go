package usecase

import (
	"context"

	"github.com/lowkruc/rataid-product-golang.git/internal/entity"
)

const (
	ErrMsgGetProducts = MsgPrefix + "GetProducts."
)

// Usecase for GetProducts
func (uc *Usecase) GetProducts(ctx context.Context, param entity.GetProductRequest) (response entity.GetProductResponse, err error) {
	// Get product from database repository
	response, err = uc.DB.GetProducts(ctx, param)
	if err != nil {
		err = wrap(err, ErrMsgGetProducts+"DB.GetProducts")
		return
	}

	return
}
