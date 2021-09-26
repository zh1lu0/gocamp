package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/zh1lu0/gocamp/internal/biz"
)

type helloRepo struct {
	data *Data
	log  *log.Helper
}

func NewHelloRepo(data *Data, logger log.Logger) biz.HelloRepo {
	return &helloRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func rdbKey() string {
	return "gocamp:rdb"
}

func (hp *helloRepo) GetRedisValue(ctx context.Context) (rv int64, err error) {
	get := hp.data.rdb.Get(ctx, rdbKey())
	rv, err = get.Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return
}

func (hp *helloRepo) IncRedisValue(ctx context.Context) error {
	_, err := hp.data.rdb.Incr(ctx, rdbKey()).Result()
	return err
}
