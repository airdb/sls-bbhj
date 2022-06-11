package mysql

import (
	"fmt"
	"sync"

	"github.com/airdb/sls-bbhj/internal/repository"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

var (
	mysqlFactory repository.Factory
	once         sync.Once
)

// GetFactoryOr create mysql factory with the given config.
func GetFactoryOr(db *gorm.DB) (repository.Factory, error) {
	var err error
	var dbIns *gorm.DB

	once.Do(func() {
		dbIns = db
		mysqlFactory = &datastore{dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %s", mysqlFactory, err.Error())
	}

	return mysqlFactory, nil
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

func (ds *datastore) Categories() repository.CategoryStore {
	return newCategory(ds)
}

func (ds *datastore) Losts() repository.LostStore {
	return newLost(ds)
}

func (ds *datastore) Files() repository.FileStore {
	return newFile(ds)
}

func (ds *datastore) Rescues() repository.RescueStore {
	return newRescue(ds)
}

func (ds *datastore) Westores() repository.WestoreStore {
	return newWestore(ds)
}
