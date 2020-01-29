package main

import (
	"github.com/airdb/mina-api/web"
	"github.com/airdb/sailor/dbutils"
)

func main() {
	dbutils.InitDefault()
	web.Run()
}
