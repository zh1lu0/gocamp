package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/zh1lu0/gocamp/internal/conf"
)

var ProviderSet = wire.NewSet(NewData, NewHelloRepo)

type Data struct {
	rdb *redis.Client
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	d := &Data{
		rdb: rdb,
	}
	cleanup := func() {
		log.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}
