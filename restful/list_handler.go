package restful

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/gy-kim/search-service/internal/data"
)

// ListService loads products
//go:generate mockery -name=ListService -case underscore -testonly -inpkg -note @generated
type ListService interface {
	Do(ctx context.Context, query string, filter *data.Filter, sort *data.SortCond, from int) ([]*data.Product, error)
}

// NewListHandler is the HTTP Handler
func NewListHandler(cfg Config, service ListService) *ListHandler {
	return &ListHandler{
		service: service,
	}
}

// ListHandler is search Product lsit
type ListHandler struct {
	cfg     Config
	service ListService
}

func (h *ListHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}

func (h *ListHandler) toJSON(writer io.Writer, products []*data.Product, page int) error {
	return errors.New("not implement")
}

type listResponseViewModel struct {
	Products []*listResponseProduct `json:"products"`
	Page     int                    `json:"page"`
}

type listResponseProduct struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Brand string `json:"brand"`
	Name  string `json:"name"`
}
