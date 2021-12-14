package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

// GetEtcdClient 获取Etcd客户端
func GetEtcdClient(cfg clientv3.Config) *clientv3.Client {
	cli, err := clientv3.New(cfg)
	if err != nil {
		panic(err)
	}
	return cli
}
