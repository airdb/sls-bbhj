package repository

import (
	"context"

	"github.com/airdb/sls-bbhj/pkg/schema"
)

// Factory defines the storage interface.
type Factory interface {
	Categories() CategoryStore
	Losts() LostStore
	Rescues() RescueStore
	Close() error
}

// LostStore defines the lost storage interface.
type LostStore interface {
	List(ctx context.Context, opts schema.LostListRequest) ([]*schema.Lost, error)
	GetByID(ctx context.Context, id int) (*schema.Lost, error)
	GetByUUID(ctx context.Context, uuid string) (*schema.Lost, error)
}

// RescueStore defines the rescue storage interface.
type RescueStore interface {
	List(ctx context.Context, opts schema.RescueListRequest) ([]*schema.Rescue, error)
}

// CategoryStore defines the lost category interface.
type CategoryStore interface {
	List(ctx context.Context, opts schema.CategoryListRequest) ([]*schema.Category, error)
	GetById(ctx context.Context, id int) (*schema.Category, error)
}
