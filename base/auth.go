package base

import "github.com/jormin/etcd-demo/pkg/helper"

// AuthStatus 获取鉴权开启状态
func AuthStatus() error {
	res, err := cli.AuthStatus(ctx)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// AuthDisable 关闭鉴权
func AuthDisable() error {
	res, err := cli.AuthDisable(ctx)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// AuthEnable 启动鉴权
func AuthEnable() error {
	res, err := cli.AuthEnable(ctx)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}
