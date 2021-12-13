package etcd

import (
	"time"

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

// GetEtcdClient 获取Etcd客户端
func GetEtcdClient() *clientv3.Client {
	cli, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	return cli
}

// GetConfig 获取配置
func GetConfig() clientv3.Config {
	return config
}
