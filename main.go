package main

import (
	"sync"
	"time"

	"github.com/airdb/mina-api/model/bo"
	"github.com/airdb/mina-api/model/po"
	"github.com/airdb/mina-api/web"
)

func main() {
	po.InitDB()

	go func() {
		wg := sync.WaitGroup{}
		interval := 300
		queueLen := 1

		for {
			wg.Add(queueLen)
			bo.SyncFrombbs(&wg)
			wg.Wait()
			<-time.After(time.Duration(interval) * time.Second)
		}
	}()

	web.Run()
}
