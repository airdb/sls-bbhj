package aggregate

import (
	"github.com/airdb/sls-bbhj/internal/repository"
)

const (
	defaultTimeFormat = "2006-01-02 15:04:05"
)

// Aggregate defines functions used to return resource interface.
type Aggregate interface {
	Lbs() LbsAggr
	Losts() LostAggr
	Westores() WestoreAggr
	Redis() *redisAggr
	Passport() *passportAggr
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

func (aggr *aggregate) Lbs() LbsAggr {
	return newLbs()
}

func (aggr *aggregate) Losts() LostAggr {
	return newLosts(aggr)
}

func (aggr *aggregate) Westores() WestoreAggr {
	return newWestores(aggr)
}

func (aggr *aggregate) Redis() *redisAggr {
	return newRedis()
}

func (aggr *aggregate) Passport() *passportAggr {
	return newPassport(aggr)
}
