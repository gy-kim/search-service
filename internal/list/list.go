package list

import (
	"context"

	"github.com/gy-kim/search-service/internal/data"
	"github.com/gy-kim/search-service/logging"
)

// Lister will attempt to load products
type Lister struct {
	cfg Config
	dao listDAO
}

// NewLister creates and initialize lister object.
func NewLister(cfg Config) *Lister {
	return &Lister{
		cfg: cfg,
	}
}

// Do load Product list
func (l *Lister) Do(ctx context.Context, query string, filter *data.Filter, sort *data.SortCond, from int) ([]*data.Product, error) {
	products, err := l.getDAO().GetProducts(ctx, query, filter, sort, from)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (l *Lister) getDAO() listDAO {
	if l.dao == nil {
		l.dao = data.NewDAO(l.cfg)
	}
	return l.dao
}

// Config is the configuration of list package.
type Config interface {
	Logger() logging.Logger
	DataURL() string
}

//go:generate mockery -name=listDAO -case underscore -testonly -inpkg -note @generated
type listDAO interface {
	GetProducts(ctx context.Context, query string, filter *data.Filter, sort *data.SortCond, from int) ([]*data.Product, error)
}
