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

// InsertProduct inserts the product to elastic.
func (d *DAO) InsertProduct(ctx context.Context, product *Product) error {
	client, err := getClient(d.cfg)
	if err != nil {
		d.logger().Error("failed to get elastic client. err: %s", err)
		return err
	}

	result, err := client.Index().
		Index(indexName).
		Type(typeName).
		Id(product.ID).
		BodyJson(product).
		Do(ctx)

	if err != nil {
		d.logger().Error("failed to insert elastic. err:%s", err)
		return err
	}

	d.logger().Info("indexed product %s to index %s type %s\n", result.Id, result.Index, result.Type)

	return nil

}

// GetProducts gets products
func (d *DAO) GetProducts(ctx context.Context, queryKeyword string, filter *Filter, sort *SortCond, page int) ([]*Product, error) {
	client, err := getClient(d.cfg)
	if err != nil {
		d.logger().Error("failed to get elastic client. err: %s", err)
		return nil, err
	}

	search := client.Search().Index(indexName).Type(typeName)

	q := elastic.NewBoolQuery()

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
	q = q.Must(queries...)
	search = search.Query(q)

	if sort != nil {
		search = search.Sort(sort.Target, sort.Ascending)
	}

	if page != 0 {
		search = search.From(page * searchSize)
	}

	search = search.Size(searchSize)

	result, err := search.Do(ctx)
	if err != nil {
		d.logger().Error("failed elastic Do. err: %s", err)
		return nil, err
	}

	d.logger().Debug("totalHists: %v", result.TotalHits())

	products, err := populateProduct(result)
	if err != nil {
		d.logger().Error("failed populate product. err: %s", err)
		return nil, err
	}

	return products, nil
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

type Filter map[string]string

type SortCond struct {
	Target    string
	Ascending bool
}

func (d *DAO) logger() logging.Logger { return d.cfg.Logger() }
