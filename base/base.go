package base

import (
	"context"
	"time"

	"github.com/jormin/etcd-demo/pkg/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// endpoints Endpoint列表
var endpoints = []string{
	"127.0.0.1:12379",
	"127.0.0.1:22379",
	"127.0.0.1:32379",
}

// config Etcd 配置
var config = clientv3.Config{
	Endpoints:            endpoints,
	DialTimeout:          time.Second * 30,
	DialKeepAliveTimeout: time.Second * 30,
	Username:             "root",
	Password:             "111111",
}
// cli Etcd客户端
var cli *clientv3.Client

// ctx 上下文
var ctx context.Context

func init() {
	cli = etcd.GetEtcdClient(config)
	ctx = context.Background()
}
