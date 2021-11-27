package aggregate

import (
	"context"
	"log"
	"strconv"

	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/airdb/sls-bbhj/pkg/schema"
)

// LostAggr defines functions used to handle lost request.
type LostAggr interface {
	List(ctx context.Context, opts schema.LostListRequest) ([]*schema.Lost, error)
}

type lostAggr struct {
	repo repository.Factory
}

var _ LostAggr = (*lostAggr)(nil)

func newLosts(aggr *aggregate) *lostAggr {
	return &lostAggr{repo: aggr.repo}
}

// List returns lost list in the storage. This function has a good performance.
func (u *lostAggr) List(ctx context.Context, opts schema.LostListRequest) ([]*schema.Lost, error) {
	if categoryId, err := strconv.Atoi(opts.Category); err == nil {
		category, err := u.repo.Categories().GetById(ctx, categoryId)
		if err == nil {
			opts.Category = category.Name
		}
	}

	losts, err := u.repo.Losts().List(ctx, opts)
	if err != nil {
		log.Printf("list losts from storage failed: %s", err.Error())

		return nil, err
	}

	return losts, nil
}
