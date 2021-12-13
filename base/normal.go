package base

import (
	"github.com/jormin/etcd-demo/pkg/helper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Put 设置值
func Put(key, val string, opts ...clientv3.OpOption) error {
	res, err := cli.Put(ctx, key, val, opts...)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// Get 获取值
func Get(key string, opts ...clientv3.OpOption) error {
	res, err := cli.Get(ctx, key, opts...)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// Delete 删除键值对
func Delete(key string, opts ...clientv3.OpOption) error {
	res, err := cli.Delete(ctx, key, opts...)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// Watch 监听
func Watch(key string, opts ...clientv3.OpOption) {
	ch := cli.Watch(ctx, key, opts...)
	for v := range ch {
		for _, val := range v.Events {
			helper.PrintJSON(val)
		}
	}
}
