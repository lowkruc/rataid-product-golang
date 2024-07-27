package usecase

import (
	"context"

	"github.com/lowkruc/rataid-product-golang.git/internal/entity"
	"github.com/lowkruc/rataid-product-golang.git/internal/repo"
	"github.com/pkg/errors"
)

const (
	MsgPrefix = "Usecase." //prefix error message
)

var (
	wrap = errors.Wrap
)

type Usecase struct {
	DB repo.DBInterface
}

// New is function to initial Usecase struct
func New(db repo.DBRepo) Usecase {
	return Usecase{
		DB: &db,
	}
}

// UsecaseInterface registered method for Usecase
type UsecaseInterface interface {
	GetProducts(ctx context.Context, param entity.GetProductRequest) (response entity.GetProductResponse, err error)
}
