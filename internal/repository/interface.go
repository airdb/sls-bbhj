package repository

import (
	"context"

	"github.com/airdb/sls-mina/pkg/schema"
)

// Factory defines the storage interface.
type Factory interface {
	Losts() LostStore
	Close() error
}

// TalkStore defines the talk storage interface.
type LostStore interface {
	List(ctx context.Context, opts schema.LostListReq) ([]*schema.Lost, error)
	GetByUUID(ctx context.Context, uuid string) (*schema.Lost, error)
}
