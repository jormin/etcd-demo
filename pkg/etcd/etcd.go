package etcd

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// Endpoints Endpoint列表
var Endpoints = []string{
	"127.0.0.1:12379",
	"127.0.0.1:22379",
	"127.0.0.1:32379",
}

// GetEtcdClient 获取Etcd客户端
func GetEtcdClient() *clientv3.Client {
	cli, err := clientv3.New(
		clientv3.Config{
			Endpoints:   Endpoints,
			DialTimeout: time.Second * 5,
			Username:    "root",
			Password:    "111111",
		},
	)
	if err != nil {
		panic(err)
	}
	return cli
}
