package repository

import (
	"context"

	"github.com/airdb/sls-bbhj/pkg/schema"
)

// Factory defines the storage interface.
type Factory interface {
	Categories() CategoryStore
	Losts() LostStore
	Files() FileStore
	Rescues() RescueStore
	Westores() WestoreStore
	Close() error
}

// LostStore defines the lost storage interface.
type LostStore interface {
	List(ctx context.Context, opts schema.LostListRequest) ([]*schema.Lost, error)
	GetByID(ctx context.Context, id uint) (*schema.Lost, error)
	GetByUUID(ctx context.Context, uuid string) (*schema.Lost, error)
	GetStatByID(ctx context.Context, id uint) (*schema.LostStat, error)
	IncreaseShare(ctx context.Context, id uint) error
	IncreaseShow(ctx context.Context, id uint) error
	Create(ctx context.Context, in schema.LostCreateRequest) error
}

// FileStore defines the lost storage interface.
type FileStore interface {
	GetLostByID(ctx context.Context, t uint) ([]*schema.File, error)
}

// RescueStore defines the rescue storage interface.
type RescueStore interface {
	List(ctx context.Context, opts schema.RescueListRequest) ([]*schema.Rescue, error)
	Count(ctx context.Context) (int64, error)
}

// WestoreStore defines the rescue storage interface.
type WestoreStore interface {
	List(ctx context.Context, opts schema.WestoreListRequest) ([]*schema.Westore, error)
	CreateOrUpdate(ctx context.Context, item *schema.Westore) error
}

// CategoryStore defines the lost category interface.
type CategoryStore interface {
	List(ctx context.Context, opts schema.CategoryListRequest) ([]*schema.Category, error)
	GetById(ctx context.Context, id int) (*schema.Category, error)
}
