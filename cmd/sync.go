package main

import (
	"sync"
	"time"

	"github.com/airdb/mina-api/model/bo"
	"github.com/airdb/mina-api/model/po"
)

func main() {
	po.InitDB()

	wg := sync.WaitGroup{}
	interval := 300
	queueLen := 1

	for {
		wg.Add(queueLen)
		// go bo.SyncAWSRoute53(&wg)
		bo.SyncFrombbs(&wg)
		wg.Wait()
		<-time.After(time.Duration(interval) * time.Second)
	}
}
