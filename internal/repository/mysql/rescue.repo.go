package mysql

import (
	"context"
	"errors"
	"log"

	"github.com/airdb/sls-mina/pkg/schema"
	"gorm.io/gorm"
)

type resuce struct {
	db *gorm.DB
}

func newRescue(ds *datastore) *resuce {
	return &resuce{db: ds.db}
}

// Create creates a new talk item.
func (r *resuce) List(ctx context.Context, opts schema.RescueListReq) ([]*schema.Rescue, error) {
	var (
		items []*schema.Rescue
		cnt   int64
	)

	tx := r.db.
		Offset(opts.PageSize * (opts.PageNo - 1)).
		Limit(opts.PageSize).
		Order("id desc")

	if len(opts.Keyword) > 0 {
		tx = tx.Where(
			"(name LIKE ? OR province LIKE ? OR city LIKE ? OR address LIKE ?)",
			"%"+opts.Keyword+"%",
			"%"+opts.Keyword+"%",
			"%"+opts.Keyword+"%",
			"%"+opts.Keyword+"%",
		)
	}

	d := tx.Find(&items).
		Offset(-1).
		Limit(-1).
		Count(&cnt)

	log.Println("len: ", len(items))

	return items, d.Error
}

// Get gets a new talk item.
func (r *resuce) GetByUUID(ctx context.Context, uuid string) (*schema.Rescue, error) {
	item := &schema.Rescue{}
	err := r.db.Where("uuid = ?", uuid).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not exist")
		}

		return nil, errors.New("can not found record")
	}

	return item, nil
}
