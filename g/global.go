package g

import (
	"context"
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
}
