package handler

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/p12s/2gis-catalog-api/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

// NewHandler - конструируем хендлер передачей в него сервиса
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitHandler() {
	http.Handle("/uppercase", httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	))

}

func ret()
uppercaseHandler := httptransport.NewServer(
makeUppercaseEndpoint(svc),
decodeUppercaseRequest,
encodeResponse,
)