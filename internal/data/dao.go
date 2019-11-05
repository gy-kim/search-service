package data

import (
	"context"
	"encoding/json"

	"github.com/gy-kim/search-service/logging"
	"github.com/olivere/elastic"
)

const (
	indexName    = "product"
	typeName     = "_doc"
	termQueryKey = "product_type"
	searchSize   = int(5)
)

const (
	defaultTotalCount = 0
)

// NewDAO initialize the data connection and return DAO reference.
func NewDAO(cfg Config) *DAO {
	client, _ = getClient(cfg)
	createAndPopulateIndex(client)

	return &DAO{cfg: cfg}
}

// DAO is the data access object.
type DAO struct {
	cfg Config
}

// GetProducts gets products
// return products []*Product, totalCount int64, err error
func (d *DAO) GetProducts(ctx context.Context, queryKeyword string, filter *Filter, sort *SortCond, page int) ([]*Product, int64, error) {
	client, err := getClient(d.cfg)
	if err != nil {
		d.logger().Error("failed to get elastic client. err: %s", err)
		return nil, defaultTotalCount, err
	}

	search := client.Search().Index(indexName).Type(typeName)
	search = d.queryCondition(search, queryKeyword, filter)
	search = d.paginateCondition(search, page)
	search = d.sortCondition(search, sort)

	result, err := search.Do(ctx)
	if err != nil {
		d.logger().Error("failed elastic Do. err: %s", err)
		return nil, defaultTotalCount, err
	}

	products, err := populateProduct(result)
	if err != nil {
		d.logger().Error("failed populate product. err: %s", err)
		return nil, defaultTotalCount, err
	}

	totalCount := result.TotalHits()
	d.logger().Debug("totalHists: %v", result.TotalHits())

	return products, totalCount, nil
}

func (d *DAO) queryCondition(service *elastic.SearchService, queryKeyword string, filter *Filter) *elastic.SearchService {
	boolQuery := elastic.NewBoolQuery()

	var queries []elastic.Query
	if queryKeyword != "" {
		query := elastic.NewTermQuery(termQueryKey, queryKeyword)
		queries = append(queries, query)
	}

	if filter != nil {
		for key, val := range *filter {
			query := elastic.NewMultiMatchQuery(val, key)
			queries = append(queries, query)
		}
	}
	boolQuery = boolQuery.Must(queries...)
	service = service.Query(boolQuery)
	return service
}

func (d *DAO) paginateCondition(service *elastic.SearchService, page int) *elastic.SearchService {
	if page > 1 {
		service = service.From((page - 1) * searchSize)
	}
	service = service.Size(searchSize)
	return service
}

func (d *DAO) sortCondition(service *elastic.SearchService, sort *SortCond) *elastic.SearchService {
	if sort != nil {
		service = service.Sort(sort.Target, sort.Ascending)
	}
	return service
}

func populateProduct(res *elastic.SearchResult) ([]*Product, error) {
	if res == nil || res.TotalHits() == 0 {
		return nil, nil
	}

	var products []*Product
	for _, hit := range res.Hits.Hits {
		product := &Product{}
		if err := json.Unmarshal(*hit.Source, product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Filter is type of Filter condition
type Filter map[string]string

// SortCond is type of Sorting condition
type SortCond struct {
	Target    string
	Ascending bool
}

func (d *DAO) logger() logging.Logger { return d.cfg.Logger() }
