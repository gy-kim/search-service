package restful

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gy-kim/search-service/internal/data"
)

var (
	errQueryMissing = errors.New("query is missing from requeset")
)

const (
	defaultPage = int(0)
)

const (
	varPageKey   = "page"
	varFilterKey = "filter"
	varQueryKey  = "q"
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
		service: service,
	}
}

// ListHandler is search Product lsit
type ListHandler struct {
	cfg     Config
	service ListService
}

func (h *ListHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query, err := h.extractQuery(request)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	filter := h.extractFilter(request)
	page := h.extractPage(request)
	sort := h.extractSort(request)

	products, err := h.service.Do(request.Context(), query, filter, sort, page)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.toJSON(response, products, page)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
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

func (h *ListHandler) extractQuery(request *http.Request) (string, error) {
	vars := mux.Vars(request)
	query, exists := vars[varQueryKey]
	if !exists {
		return "", errQueryMissing
	}
	return query, nil
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
