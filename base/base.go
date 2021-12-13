package base

import (
	"context"

	"github.com/jormin/etcd-demo/pkg/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// cli Etcd客户端
var cli *clientv3.Client

// ctx 上下文
var ctx context.Context

func init() {
	cli = etcd.GetEtcdClient()
	ctx = context.Background()
}
