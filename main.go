package main

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/fasthall/gochariots/maintainer/adapter/mongodb"
	"github.com/fasthall/redis-cache-for-lowgo/cache"
	"github.com/fasthall/redis-cache-for-lowgo/config"
	"github.com/go-redis/redis"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var client *redis.Client
var mongoHost []string
var mongoVer int
var mongoClient []mongodb.Client

type Server struct{}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (s *Server) UpdateStorage(ctx context.Context, in *cache.RPCStorages) (*cache.RPCReply, error) {
	ver := int(in.GetVersion())
	if ver > mongoVer {
		mongoVer = ver
		mongoHost = in.GetHosts()
		mongoClient = make([]mongodb.Client, len(mongoHost))
		for i := range mongoHost {
			mongoClient[i] = mongodb.NewClient(mongoHost[i])
		}
		logrus.WithField("host", in.GetHosts()).Info("received mongoDB hosts update")
	} else {
		logrus.WithFields(logrus.Fields{"current": mongoVer, "received": ver}).Debug("received older version of mongoDB hosts")
	}
	return &cache.RPCReply{Message: "ok"}, nil
}

func (s *Server) Put(ctx context.Context, in *cache.RPCIDs) (*cache.RPCReply, error) {
	pipe := client.Pipeline()
	for _, id := range in.GetIds() {
		err := pipe.Set(id, 1, 0).Err()
		if err != nil {
			logrus.WithError(err).Println("failed to update cache")
		}
	}
	_, err := pipe.Exec()
	if err != nil {
		return &cache.RPCReply{Message: "pipeline error"}, err
	}
	return &cache.RPCReply{Message: "ok"}, nil
}

func (s *Server) Get(ctx context.Context, in *cache.RPCIDs) (*cache.RPCBools, error) {
	pipe := client.Pipeline()
	cmds := make([]*redis.StringCmd, len(in.GetIds()))
	for i, id := range in.GetIds() {
		cmds[i] = pipe.Get(id)
	}
	_, err := pipe.Exec()
	exists := make([]bool, len(cmds))
	for i, cmd := range cmds {
		if cmd.Err() != nil {
			// cache miss, fetch from storage
			exists[i] = false
		} else {
			if cmd.Val() == "1" {
				exists[i] = true
			} else {
				exists[i] = false
			}
		}
	}
	if err != nil && err != redis.Nil {
		return &cache.RPCBools{Exists: []bool{}}, err
	}
	return &cache.RPCBools{Exists: exists}, nil
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

	go func() {
		err = errors.New("")
		for err != nil {
			time.Sleep(time.Second * 1)
			_, _, err = config.Report()
			if err != nil {
				logrus.WithError(err).Error("failed to report to controller, retry in 1 second")
			}
		}
	}()

	fmt.Println("server listening to 6380")
	if err := s.Serve(ln); err != nil {
		panic(err)
	}
}
