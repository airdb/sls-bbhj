package repository

import "github.com/airdb/sls-mina/pkg/schema"

// Factory defines the storage interface.
type Factory interface {
	Losts() LostStore
	Close() error
}

// TalkStore defines the talk storage interface.
type LostStore interface {
	List() ([]*schema.Lost, error)
}
