package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Hello struct {
	RedisValue int64
}

type HelloRepo interface {
	// redis
	GetRedisValue(ctx context.Context) (rv int64, err error)
	IncRedisValue(ctx context.Context) error
}

type HelloUsecase struct {
	repo HelloRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo HelloRepo, logger log.Logger) *HelloUsecase {
	return &HelloUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *HelloUsecase) Get(ctx context.Context) (p *Hello, err error) {

	err = uc.repo.IncRedisValue(ctx)
	if err != nil {
		return
	}

	p = &Hello{}
	p.RedisValue, err = uc.repo.GetRedisValue(ctx)
	return
}
