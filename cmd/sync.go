package main

import (
	"fmt"

	"github.com/airdb/mina-api/model/po"
	"github.com/airdb/sailor/dbutils"
)

func main() {
	dbutils.InitDefault()

	aa := po.GetBBSArticles()

	fmt.Print(aa)
}
