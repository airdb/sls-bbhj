package repository

// Factory defines the storage interface.
type Factory interface {
	Losts() LostStore
	Close() error
}

// TalkStore defines the talk storage interface.
type LostStore interface {
	List() error
}
