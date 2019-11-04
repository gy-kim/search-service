package data

import (
	"context"
	"errors"

	"github.com/gy-kim/search-service/logging"
	"github.com/olivere/elastic"
)

const (
	indexName    = "product"
	typeName     = "product"
	termQueryKey = "type"
	searchSize   = int(5)
)

// NewDAO initialize the data connection and return DAO reference.
func NewDAO(cfg Config) *DAO {
	_, _ = getClient(cfg)

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

func (d *DAO) GetProducts(ctx context.Context, query string, filter *Filter, sort *SortCond, from int) ([]*Product, error) {
	client, err := getClient(d.cfg)
	if err != nil {
		d.logger().Error("failed to get elastic client. err: %s", err)
		return nil, err
	}

	termQuery := elastic.NewTermQuery(termQueryKey, query)

	searchService := client.Search().
		Index(indexName).
		Query(termQuery).
		From(from).Size(searchSize)

	if sort != nil {
		searchService = searchService.Sort(sort.Target, sort.Ascending)
	}

	return nil, errors.New("not implement")

}

type Filter map[string]string

type SortCond struct {
	Target    string
	Ascending bool
}

func (d *DAO) logger() logging.Logger { return d.cfg.Logger() }
