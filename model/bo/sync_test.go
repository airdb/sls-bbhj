package bo

import (
	"testing"

	"github.com/airdb/sailor/dbutils"
)

func TestSyncFrombbsByID(t *testing.T) {
	dbutils.InitDefault()

	var tid uint = 425750

	SyncFrombbsByID(tid)
}
