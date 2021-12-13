package base

import (
	"github.com/jormin/etcd-demo/pkg/helper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// RoleList 获取角色列表
func RoleList() error {
	res, err := cli.RoleList(ctx)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// RoleAdd 添加角色
func RoleAdd(role string) error {
	res, err := cli.RoleAdd(ctx, role)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// RoleGet 获取角色信息
func RoleGet(role string) error {
	res, err := cli.RoleGet(ctx, role)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// RoleDelete 删除角色
func RoleDelete(role string) error {
	res, err := cli.RoleDelete(ctx, role)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// RoleGrantPermission 赋予角色权限
func RoleGrantPermission(role string, key, rangeEnd string, permType clientv3.PermissionType) error {
	res, err := cli.RoleGrantPermission(ctx, role, key, rangeEnd, permType)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// RoleRevokePermission 撤销角色权限
func RoleRevokePermission(role string, key, rangeEnd string) error {
	res, err := cli.RoleRevokePermission(ctx, role, key, rangeEnd)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}
