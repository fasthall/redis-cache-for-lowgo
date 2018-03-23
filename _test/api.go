package main

import (
	"context"
	"fmt"

	"github.com/fasthall/redis-cache-for-lowgo/cache"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func main() {
	a := uuid.New().String()
	b := uuid.New().String()
	c := uuid.New().String()
	d := uuid.New().String()
	conn, err := grpc.Dial("localhost:6380", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := cache.NewCacheClient(conn)
	ids := &cache.RPCIDs{
		Ids: []string{a, b, d},
	}
	_, err = client.Put(context.Background(), ids)
	if err != nil {
		panic(err)
	}

	ids = &cache.RPCIDs{
		Ids: []string{a, b, c, d},
	}
	result, err := client.Get(context.Background(), ids)
	fmt.Println(result.GetExists(), err)
}
