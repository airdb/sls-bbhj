package aggregate

import "github.com/airdb/sls-bbhj/internal/repository"

// Aggregate defines functions used to return resource interface.
type Aggregate interface {
	Losts() LostAggr
}

type aggregate struct {
	repo repository.Factory
}

// New returns Aggregate interface.
func New(repo repository.Factory) Aggregate {
	return &aggregate{
		repo: repo,
	}
}
func (aggr *aggregate) Losts() LostAggr {
	return newLosts(aggr)
}
