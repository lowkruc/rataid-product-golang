package handler

import "net/http"

type HTTPHandler struct {
}

type HTTPInterface interface {
	GetProduct(w http.ResponseWriter, r *http.Request)
}
