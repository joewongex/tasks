package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joewongex/tasks/g"
	"github.com/joewongex/tasks/kjwj"
	"github.com/robfig/cron/v3"
)

func main() {
	g.Log.Info("程序开始执行")
	c := cron.New()
	ctx := context.Background()
	wg := sync.WaitGroup{}
	sigs := make(chan os.Signal, 1)

	c.AddFunc("@hourly", func() {
		kjwj.CheckLatestNews(ctx, wg)
	})
	c.Start()

	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-sigs:
		g.Log.Info("接收到终止信号，程序准备退出")
	}

	wg.Wait()
	g.Log.Info("程序已结束")
}
