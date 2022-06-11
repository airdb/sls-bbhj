package mysql

import (
	"context"
	"log"

	"github.com/airdb/sls-bbhj/pkg/schema"
	"gorm.io/gorm"
)

type file struct {
	db *gorm.DB
}

func newFile(ds *datastore) *file {
	return &file{db: ds.db}
}

// Create creates a new talk item.
func (r *file) GetLostByID(ctx context.Context, id uint) ([]*schema.File, error) {
	var (
		items []*schema.File
		cnt   int64
	)

	tx := r.db.
		Order("sort_id desc")

	tx = tx.Where("type = 'lost' and parent_id = ?", id)

	d := tx.Find(&items).
		Offset(-1).
		Limit(-1).
		Count(&cnt)

	log.Println("len: ", len(items))

	return items, d.Error
}
