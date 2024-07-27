package handler

import (
	"net/http"

	"github.com/lowkruc/rataid-product-golang.git/internal/common"
	"github.com/lowkruc/rataid-product-golang.git/internal/usecase"
)

var (
	sendInternalErr = common.SendInternalErr
	sendOk          = common.SendOk
	sendOkWithData  = common.SendOkWithData
	sendBadRequest  = common.SendBadRequest
)

type HTTPHandler struct {
	UC usecase.UsecaseInterface
}

// NewHTTPHandler is function to initial HTTPHandler struct
func NewHTTPHandler(uc usecase.Usecase) HTTPHandler {
	return HTTPHandler{
		UC: &uc,
	}
}

// HTTPInterface registered method for HTTP Handler
type HTTPInterface interface {
	GetProduct(w http.ResponseWriter, r *http.Request)
}
