package g

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v9"
	"github.com/joewongex/tasks/module"
	"github.com/sirupsen/logrus"
)

var (
	Config *module.Config
	Client *http.Client
	Redis  *redis.Client
	Log    *logrus.Logger
)

type logErrHook struct{}

func (h *logErrHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel}
}

func (h *logErrHook) Fire(e *logrus.Entry) error {
	msg := fmt.Sprintf("程序发生[%s]错误：%s", e.Level.String(), e.Message)
	return BarkNotify("Tasks定时任务", msg)
}

func init() {
	var err error
	Config, err = module.NewConfig()
	if err != nil {
		panic(err)
	}

	Log, err = module.NewLog(Config)
	if err != nil {
		log.Fatalf("创建日志模块失败：%v", err)
	}

	Client, err = module.NewClient()
	if err != nil {
		Log.Fatalf("创建Client模块失败：%v", err)
	}

	Redis = module.NewRedis(Config)
	err = Redis.Ping(context.Background()).Err()
	if err != nil {
		Log.Fatalf("创建Redis模块失败：%v", err)
	}

	Log.AddHook(&logErrHook{})
}
