package aggregate

import (
	"context"
	"log"
	"testing"

	"github.com/airdb/sailor/dbutil"
	"github.com/airdb/sls-bbhj/internal/app"
	"github.com/airdb/sls-bbhj/internal/repository/mysql"
)

func Test_westoreAggr_SyncLoc(t *testing.T) {
	app.InitApp()

	repo, err := mysql.GetFactoryOr(dbutil.WriteDefaultDB())
	if err != nil {
		log.Panic(err)
	}

	aggr := New(repo)

	aggr.Westores().SyncLoc(context.Background())
}
