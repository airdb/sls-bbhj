package mysql

import (
	"context"
	"errors"
	"log"

	"github.com/airdb/sls-bbhj/pkg/schema"
	"gorm.io/gorm"
)

type category struct {
	db *gorm.DB
}

func newCategory(ds *datastore) *category {
	return &category{db: ds.db}
}

// Create creates a new talk item.
func (r *category) List(ctx context.Context, opts schema.CategoryListRequest) ([]*schema.Category, error) {
	var (
		items []*schema.Category
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
func (r *category) GetById(ctx context.Context, id int) (*schema.Category, error) {
	item := &schema.Category{}
	err := r.db.Where("id = ?", id).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not exist")
		}

		return nil, errors.New("can not found record")
	}

	return item, nil
}
