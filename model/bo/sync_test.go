package bo_test

import (
	"testing"

	"github.com/airdb/mina-api/model/bo"
	"github.com/airdb/mina-api/model/po"
)

func TestSyncFrombbsByID(t *testing.T) {
	po.InitDB()

	var tid uint = 425750

	bo.SyncFrombbsByID(tid)
}
