package app

import (
	"github.com/airdb/sailor/dbutil"
)

func InitApp() {
	// Init Database.
	dbutil.InitDefaultDB()
}
