package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type HttpServer struct {
	transport string
}

func (s *HttpServer) Start() error {
	fmt.Printf("I am start at %s \n", s.transport)
	return nil
}

func (s *HttpServer) Stop() error {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("I am stoped at %s \n", s.transport)
	return nil
}

func NewHttpServer(port string) *HttpServer {
	return &HttpServer{port}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	g, gctx := errgroup.WithContext(ctx)

	s := NewHttpServer("8080")

	g.Go(func() error {
		<-ctx.Done()
		return s.Stop()
	})

	g.Go(func() error {
		return s.Start()
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	g.Go(func() error {
		for {
			select {
			case <-gctx.Done():
				return ctx.Err()
			case <-c:
				fmt.Println("Get quit signal")
				s.Stop()
				return errors.New("Terminated")
			}
		}
	})

	g.Go(func() error {
		// 模拟手动关闭系统？
		time.Sleep(5 * time.Second)
		cancel()
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server stoped")
}
