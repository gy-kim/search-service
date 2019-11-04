package data

import (
	"context"

	"github.com/gy-kim/search-service/logging"
)

const (
	indexName = "product"
	typeName  = "product"
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

func (d *DAO) logger() logging.Logger { return d.cfg.Logger() }
