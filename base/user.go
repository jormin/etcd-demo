package base

import (
	"github.com/jormin/etcd-demo/pkg/helper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// UserList 获取用户列表
func UserList() error {
	res, err := cli.UserList(ctx)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// UserAdd 添加用户
func UserAdd(name, password string) error {
	res, err := cli.UserAdd(ctx, name, password)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// UserAddWithOptions 使用选项添加用户
func UserAddWithOptions(name, password string, options *clientv3.UserAddOptions) error {
	res, err := cli.UserAddWithOptions(ctx, name, password, options)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// UserChangePasswd 修改密码
func UserChangePasswd(name, password string) error {
	res, err := cli.UserChangePassword(ctx, name, password)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// UserGet 获取用户信息
func UserGet(name string) error {
	res, err := cli.UserGet(ctx, name)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// UserDelete 删除用户
func UserDelete(name string) error {
	res, err := cli.UserDelete(ctx, name)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// UserGrantRole 赋予用户角色
func UserGrantRole(name, role string) error {
	res, err := cli.UserGrantRole(ctx, name, role)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// UserRevokeRole 撤销用户角色
func UserRevokeRole(name, role string) error {
	res, err := cli.UserRevokeRole(ctx, name, role)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}
