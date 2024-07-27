package repo

import (
	"context"

	"github.com/lowkruc/rataid-product-golang.git/internal/entity"
)

type DBRepo struct {
	// TODO: Initial pointer to database PostgreSQL or any module
}

type DBInterface interface {
	GetProducts(ctx context.Context, param entity.GetProductRequest) (result entity.GetProductResponse, err error)
}
