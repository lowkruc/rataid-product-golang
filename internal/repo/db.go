package repo

import (
	"context"

	"github.com/lowkruc/rataid-product-golang.git/internal/entity"
)

type DBRepo struct {
	// TODO: Initial pointer to database PostgreSQL or any module
}

// New is function to initial DBRepo struct
func New() DBRepo {
	return DBRepo{}
}

// DBInterface registered method for Database repository
type DBInterface interface {
	GetProducts(ctx context.Context, param entity.GetProductRequest) (result entity.GetProductResponse, err error)
}
