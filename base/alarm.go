package base

import (
	"github.com/jormin/etcd-demo/pkg/helper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// AlarmList 获取警报列表
func AlarmList() error {
	res, err := cli.AlarmList(ctx)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// AlarmDisarm 解除警报
func AlarmDisarm(m *clientv3.AlarmMember) error {
	res, err := cli.AlarmDisarm(ctx, m)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}
