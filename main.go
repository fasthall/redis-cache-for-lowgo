package main

import (
	"github.com/fasthall/redis-cache-for-lowgo/cache"
	context "golang.org/x/net/context"
)

type Server struct{}

func (s *Server) Put(ctx context.Context, in *cache.RPCIDs) (*cache.RPCReply, error) {
	return &cache.RPCReply{Message: "ok"}, nil
}

func (s *Server) Get(ctx context.Context, in *cache.RPCIDs) (*cache.RPCBools, error) {
	return &cache.RPCBools{Exists: []bool{}}, nil
}

func main() {

}
