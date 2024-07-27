package usecase

import "github.com/lowkruc/rataid-product-golang.git/internal/repo"

type Usecase struct {
	DB *repo.DBInterface
}
