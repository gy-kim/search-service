package restful

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gy-kim/search-service/internal/data"
	"github.com/gy-kim/search-service/logging"
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

// serveHTTP implement http.Handler
// http://127.0.0.1:9000/v1/products?q=black_shoes&filter=brand:adidas&page=2&sort=name&sort_asc=false
func (h *ListHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	query := h.extractQuery(request)
	filter := h.extractFilter(request)
	page := h.extractPage(request)
	sort := h.extractSort(request)

	products, err := h.service.Do(ctx, query, filter, sort, page)
	if err != nil {
		h.logger().Error("Failed to Do. err:%s", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.writeJSON(response, products, page)
	if err != nil {
		h.logger().Error("Failed toJSON. err:%s", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *ListHandler) extractSort(r *http.Request) *data.SortCond {
	target, exists := r.URL.Query()[varSortKey]
	if !exists || len(target) < 1 {
		return nil
	}

	asc := true
	str, exists := r.URL.Query()[varSortAsc]
	if exists && len(str) >= 1 {
		if str[0] == "false" {
			asc = false
		}
	}

	sort := &data.SortCond{
		Target:    target[0],
		Ascending: asc,
	}
	h.logger().Debug("[extractSort] sort: (%#v)", sort)
	return sort
}

func (h *ListHandler) extractPage(r *http.Request) int {
	str, exists := r.URL.Query()[varPageKey]
	if !exists || len(str) < 1 {
		return defaultPage
	}
	page, err := strconv.Atoi(str[0])
	if err != nil {
		return defaultPage
	}
	h.logger().Debug("[extractPage] Page:(%v)", page)
	return page
}

func (h *ListHandler) extractFilter(r *http.Request) *data.Filter {
	str, exists := r.URL.Query()[varFilterKey]
	if !exists || len(str) < 1 {
		return nil
	}

	arr := strings.Split(str[0], ":")
	if len(arr) != 2 {
		return nil
	}
	filter := &data.Filter{arr[0]: arr[1]}
	h.logger().Debug("[extractFilter] filter:(%v)", filter)
	return filter
}

func (h *ListHandler) extractQuery(r *http.Request) string {
	query, exists := r.URL.Query()[varQueryKey]
	if !exists || len(query) < 1 {
		return ""
	}
	h.logger().Debug("[extractQuery] query:(%s)", query[0])
	return query[0]
}

func (h *ListHandler) writeJSON(writer io.Writer, products []*data.Product, page int) error {
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
