package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/zh1lu0/gocamp/api/week4/v1"
	"github.com/zh1lu0/gocamp/internal/biz"
)

type HelloService struct {
	pb.UnimplementedHelloServer

	uc  *biz.HelloUsecase
	log *log.Helper
}

func NewHelloService(uc *biz.HelloUsecase, logger log.Logger) *HelloService {
	return &HelloService{uc: uc, log: log.NewHelper(logger)}
}

func (s *HelloService) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", req.GetName())

	p, err := s.uc.Get(ctx)
	if err != nil {
		return nil, err
	}
	s.log.WithContext(ctx).Infof("SayHello Current Redis value: %v", p.RedisValue)

	return &pb.SayHelloReply{Message: fmt.Sprintf("Hello %v, internal value is : %v", req.GetName(), p.RedisValue)}, nil
}
