package base

import (
	"github.com/jormin/etcd-demo/pkg/helper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// LeaseList 获取租约列表
func LeaseList() error {
	res, err := cli.Leases(ctx)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// LeaseGrant 生成租约
func LeaseGrant(ttl int64) error {
	res, err := cli.Grant(ctx, ttl)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// LeaseRevoke 撤销租约
func LeaseRevoke(id int64) error {
	res, err := cli.Revoke(ctx, clientv3.LeaseID(id))
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// LeaseTimeToLive 租约信息
func LeaseTimeToLive(id int64, opts ...clientv3.LeaseOption) error {
	res, err := cli.TimeToLive(ctx, clientv3.LeaseID(id), opts...)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// LeaseKeepAlive 保持租约不过期
func LeaseKeepAlive(id int64) error {
	res, err := cli.KeepAlive(ctx, clientv3.LeaseID(id))
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// LeaseKeepAliveOnce 更新租约一次
func LeaseKeepAliveOnce(id int64) error {
	res, err := cli.KeepAliveOnce(ctx, clientv3.LeaseID(id))
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}
