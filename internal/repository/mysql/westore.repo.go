package mysql

import (
	"context"
	"errors"
	"log"

	"github.com/airdb/sls-bbhj/pkg/schema"
	"gorm.io/gorm"
)

type westore struct {
	db *gorm.DB
}

func newWestore(ds *datastore) *westore {
	return &westore{db: ds.db}
}

func (r *westore) CreateOrUpdate(ctx context.Context, item *schema.Westore) error {
	var (
		cnt int64
		err error
	)

	err = r.db.Model(item).Where(schema.Westore{Address: item.Address}).Count(&cnt).Error
	if err != nil {
		return err
	}
	if cnt == 0 {
		err = r.db.Create(item).Error
		return err
	}

	err = r.db.Model(item).
		Where(schema.Westore{Address: item.Address}).
		Updates(schema.Westore{
			Longitude: item.Longitude,
			Latitude:  item.Latitude,
		}).
		Error
	if err != nil {
		return err
	}

	err = r.db.Model(item).Where(schema.Westore{Address: item.Address}).
		First(&item).Error
	if err != nil {
		return err
	}

	return err
}

// Create creates a new talk item.
func (r *westore) List(ctx context.Context, opts schema.WestoreListRequest) ([]*schema.Westore, error) {
	var (
		items []*schema.Westore
		cnt   int64
	)

	tx := r.db.
		Offset(opts.PageSize * (opts.PageNo - 1)).
		Limit(opts.PageSize).
		Order("id desc")

	if len(opts.Keyword) > 0 {
		queryWord := "%" + opts.Keyword + "%"
		tx = tx.Where("(nickname like ?) OR (missed_address like ?)",
			queryWord,
			queryWord,
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
func (r *westore) GetByID(ctx context.Context, id uint) (*schema.Westore, error) {
	item := &schema.Westore{}
	err := r.db.Where("id = ?", id).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not exist")
		}

		return nil, errors.New("can not found record")
	}

	return item, nil
}

// Get gets a new talk item.
func (r *westore) GetByUUID(ctx context.Context, uuid string) (*schema.Westore, error) {
	item := &schema.Westore{}
	err := r.db.Where("uuid = ?", uuid).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not exist")
		}

		return nil, errors.New("can not found record")
	}

	return item, nil
}
