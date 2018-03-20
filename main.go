package main

import (
	"fmt"
	"net"

	"github.com/fasthall/redis-cache-for-lowgo/cache"
	"github.com/fasthall/redis-cache-for-lowgo/config"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct{}

func (s *Server) Put(ctx context.Context, in *cache.RPCIDs) (*cache.RPCReply, error) {
	return &cache.RPCReply{Message: "ok"}, nil
}

func (s *Server) Get(ctx context.Context, in *cache.RPCIDs) (*cache.RPCBools, error) {
	return &cache.RPCBools{Exists: []bool{}}, nil
}

func main() {
	s := grpc.NewServer()
	cache.RegisterCacheServer(s, &Server{})
	reflection.Register(s)
	ln, err := net.Listen("tcp", ":6380")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("server listening to 6380")
	_, _, err = config.Report()
	if err != nil {
		panic(err)
	}
	if err := s.Serve(ln); err != nil {
		panic(err)
	}
}
