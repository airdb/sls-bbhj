package aggregate

import (
	"context"
	"log"
	"time"

	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/airdb/sls-bbhj/pkg/schema"
)

// WestoreAggr defines functions used to handle westore request.
type WestoreAggr interface {
	SyncLoc(ctx context.Context) error
}

type westoreAggr struct {
	repo repository.Factory
	aggr aggregate
}

var _ WestoreAggr = (*westoreAggr)(nil)

func newWestores(aggr *aggregate) *westoreAggr {
	return &westoreAggr{repo: aggr.repo}
}

// List returns westore list in the storage. This function has a good performance.
func (u *westoreAggr) SyncLoc(ctx context.Context) error {
	total, err := u.repo.Rescues().Count(ctx)
	if err != nil {
		log.Printf("list westores from storage failed: %s", err.Error())

		return err
	}

	opt := schema.RescueListRequest{}
	opt.PageSize = 100
	opt.PageNo = 0
	for int64(opt.PageNo*opt.PageSize) < total {
		opt.PageNo += 1
		items, err := u.repo.Rescues().List(ctx, opt)
		if err != nil {
			log.Printf("list westores from storage failed: %s", err.Error())

			return err
		}

		for idx, item := range items {
			resp, err := u.aggr.Lbs().GeoCoder(ctx, item.Address)
			if err != nil {
				continue
			}
			ws := schema.Westore{
				Name:      item.Name,
				Longitude: resp.Result.Location.Lng,
				Latitude:  resp.Result.Location.Lat,
				Address:   item.Address,
				Province:  resp.Result.AddressComponents.Province,
				City:      resp.Result.AddressComponents.City,
				District:  resp.Result.AddressComponents.District,
				Phone1:    item.Phone1,
				Phone2:    item.Phone2,
			}
			u.repo.Westores().CreateOrUpdate(ctx, &ws)

			if idx%5 == 0 {
				time.Sleep(time.Second + time.Microsecond*100)
			}
		}

	}

	return nil
}
