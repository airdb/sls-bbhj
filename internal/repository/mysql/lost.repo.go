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
func (r *lost) List(ctx context.Context, opts schema.LostListRequest) ([]*schema.Lost, error) {
	var (
		items []*schema.Lost
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

	if len(opts.Category) > 0 {
		tx = tx.Where("category = ?", opts.Category)
	}

	d := tx.Find(&items).
		Offset(-1).
		Limit(-1).
		Count(&cnt)

	log.Println("len: ", len(items))

	return items, d.Error
}

// Get gets a new talk item.
func (r *lost) GetByID(ctx context.Context, id uint) (*schema.Lost, error) {
	item := &schema.Lost{}
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

// CreateStatByID crate a lost stat from a lost.
func (r *lost) CreateStatByID(ctx context.Context, id uint) (*schema.LostStat, error) {
	lost, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	item := &schema.LostStat{
		LostID: lost.ID,
		Babyid: lost.Babyid,
	}
	err = r.db.Create(item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not exist")
		}

		return nil, errors.New("can not found record")
	}

	return item, nil
}

// GetStatByID get a lost stat by a lost id. if not exist then create it.
func (r *lost) GetStatByID(ctx context.Context, id uint) (*schema.LostStat, error) {
	item := &schema.LostStat{}
	err := r.db.Where("lost_id = ?", id).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r.CreateStatByID(ctx, id)
		}

		return nil, errors.New("can not found record")
	}

	return item, nil
}

// IncreaseShare get a lost stat by a lost id. if not exist then create it.
func (r *lost) IncreaseShare(ctx context.Context, id uint) error {
	item := &schema.LostStat{
		LostID: id,
	}
	err := r.db.Model(item).
		Where("lost_id = ?", id).
		UpdateColumn("share_count", gorm.Expr("share_count + ?", 1)).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		return errors.New("can not increse share")
	}

	return nil
}

// IncreaseShow get a lost stat by a lost id. if not exist then create it.
func (r *lost) IncreaseShow(ctx context.Context, id uint) error {
	item := &schema.LostStat{
		LostID: id,
	}
	err := r.db.Model(item).
		Where("lost_id = ?", id).
		UpdateColumn("show_count", gorm.Expr("show_count + ?", 1)).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		return errors.New("can not increse show")
	}

	return nil
}
