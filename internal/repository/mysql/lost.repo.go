package mysql

import (
	"context"
	"errors"
	"log"

	"github.com/airdb/sls-bbhj/pkg/schema"
	"gorm.io/gorm"
)

type lost struct {
	db *gorm.DB
}

func newLost(ds *datastore) *lost {
	return &lost{db: ds.db}
}

// Create creates a new talk item.
func (r *lost) List(ctx context.Context, opts schema.LostListReq) ([]*schema.Lost, error) {
	var (
		items []*schema.Lost
		cnt   int64
	)

	tx := r.db.
		Offset(opts.PageSize * (opts.PageNo - 1)).
		Limit(opts.PageSize).
		Order("id desc")

	if len(opts.Keyword) > 0 {
		tx = tx.Where("name like ?", "%"+opts.Keyword+"%")
	}

	d := tx.Find(&items).
		Offset(-1).
		Limit(-1).
		Count(&cnt)

	log.Println("len: ", len(items))

	return items, d.Error
}

// Get gets a new talk item.
func (r *lost) GetByUUID(ctx context.Context, uuid string) (*schema.Lost, error) {
	item := &schema.Lost{}
	err := r.db.Where("uuid = ?", uuid).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not exist")
		}

		return nil, errors.New("can not found record")
	}

	return item, nil
}
