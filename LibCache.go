package golib_v1

import (
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
)

var (
	ctx    = context.Background()
	client = &redisClusterClient{}
)

type redisClusterClient struct {
	clusterClient *redis.ClusterClient
}

func InitializeRedisCluster(hostnames string) *redisClusterClient {

	address := strings.Split(hostnames, ",")

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: address,
	})

	rdb.Ping(ctx)
	client.clusterClient = rdb
	return client
}

func RedisSentinal(hostnames string) *redis.ClusterClient {
	address := strings.Split(hostnames, ",")

	rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:    "master-name",
		SentinelAddrs: address,
	})

	return rdb

}

func InitializeRedisSentinal(address string) {

	sentinel := redis.NewSentinelClient(&redis.Options{
		Addr: address, //:9126
	})

	sentinel.GetMasterAddrByName(ctx, "master-name").Result()

}
