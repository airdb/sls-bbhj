package repository

import (
	"context"

	"github.com/airdb/sls-bbhj/pkg/schema"
)

// Factory defines the storage interface.
type Factory interface {
	Losts() LostStore
	Rescues() RescueStore
	Close() error
}

// LostStore defines the lost storage interface.
type LostStore interface {
	List(ctx context.Context, opts schema.LostListReq) ([]*schema.Lost, error)
	GetByID(ctx context.Context, id int) (*schema.Lost, error)
	GetByUUID(ctx context.Context, uuid string) (*schema.Lost, error)
}

// RescueStore defines the talk storage interface.
type RescueStore interface {
	List(ctx context.Context, opts schema.RescueListReq) ([]*schema.Rescue, error)
}
