package handler

import (
	"net/http"
	"strconv"

	"github.com/lowkruc/rataid-product-golang.git/internal/entity"
)

// GetProduct is func to handle request API `GET /v1/product`
func (h *HTTPHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// var context from request
	ctx := r.Context()

	// initial request parameter product
	param := entity.GetProductRequest{
		Limit: 10,
		Page:  1,
	}

	// TODO: create function to handle binding query param and json body data
	limitParam := r.URL.Query().Get("limit")
	if limitParam != "" {
		limitInt, err := strconv.Atoi(limitParam)
		if err != nil {
			sendBadRequest(w, err.Error())
		}

		param.Limit = limitInt
	}

	pageParam := r.URL.Query().Get("page")
	if pageParam != "" {
		pageInt, err := strconv.Atoi(pageParam)
		if err != nil {
			sendBadRequest(w, err.Error())
			return
		}

		param.Page = pageInt
	}

	// get product from usecase
	response, err := h.UC.GetProducts(ctx, param)

	// if err send internal server error
	// TODO: send slack error notify
	if err != nil {
		sendInternalErr(w)
		return
	}

	// send data response
	sendOkWithData(w, response)
}
