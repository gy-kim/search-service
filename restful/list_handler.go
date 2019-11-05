package restful

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gy-kim/search-service/internal/data"
	"github.com/gy-kim/search-service/logging"
)

var (
	errQueryMissing = errors.New("query is missing from requeset")
)

const (
	defaultPage = int(0)
)

const (
	varQueryKey  = "q"
	varPageKey   = "page"
	varFilterKey = "filter"
	varSortKey   = "sort"
	varSortAsc   = "sort_asc"
)

// ListService loads products
//go:generate mockery -name=ListService -case underscore -testonly -inpkg -note @generated
type ListService interface {
	Do(ctx context.Context, query string, filter *data.Filter, sort *data.SortCond, from int) ([]*data.Product, error)
}

// NewListHandler is the HTTP Handler
func NewListHandler(cfg Config, service ListService) *ListHandler {
	return &ListHandler{
		cfg:     cfg,
		service: service,
	}
}

// ListHandler is search Product lsit
type ListHandler struct {
	cfg     Config
	service ListService
}

func (h *ListHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	h.logger().Debug("ServeHTTP")
	query, err := h.extractQuery(request)
	if err != nil {
		h.logger().Error("Failed to extract query. err:%s", err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	filter := h.extractFilter(request)
	page := h.extractPage(request)
	sort := h.extractSort(request)

	products, err := h.service.Do(request.Context(), query, filter, sort, page)
	if err != nil {
		h.logger().Error("Failed to Do. err:%s", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.toJSON(response, products, page)
	if err != nil {
		h.logger().Error("Failed toJSON. err:%s", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *ListHandler) extractSort(request *http.Request) *data.SortCond {
	vars := mux.Vars(request)
	target, exists := vars[varSortKey]
	if !exists {
		return nil
	}

	asc := true

	if str, exists := vars[varSortAsc]; exists {
		if str == "false" {
			asc = false
		}
	}

	sort := &data.SortCond{
		Target:    target,
		Ascending: asc,
	}
	return sort
}

func (h *ListHandler) extractPage(request *http.Request) int {
	vars := mux.Vars(request)
	str, exists := vars[varPageKey]
	if !exists {
		return defaultPage
	}
	page, err := strconv.Atoi(str)
	if err != nil {
		return defaultPage
	}
	return page
}

func (h *ListHandler) extractFilter(request *http.Request) *data.Filter {
	vars := mux.Vars(request)
	str, exists := vars[varFilterKey]
	if !exists {
		return nil
	}

	arr := strings.Split(str, ":")
	if len(arr) != 2 {
		return nil
	}
	filter := &data.Filter{arr[0]: arr[1]}
	return filter
}

func (h *ListHandler) extractQuery(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	query, exists := vars[varQueryKey]
	if !exists {
		return "", errQueryMissing
	}
	return query, nil
}

func (h *ListHandler) toJSON(writer io.Writer, products []*data.Product, page int) error {
	out := &listResponseViewModel{
		Page:     page,
		Products: make([]*listResponseProduct, len(products)),
	}
	for idx, p := range products {
		out.Products[idx] = &listResponseProduct{
			ID:    p.ID,
			Type:  p.Type,
			Brand: p.Brand,
			Name:  p.Name,
		}
	}

	return json.NewEncoder(writer).Encode(out)
}

func (h *ListHandler) logger() logging.Logger { return h.cfg.Logger() }

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
