package base

import (
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// TestRoleList 测试获取角色列表
func TestRoleList(t *testing.T) {
	if err := RoleList(); err != nil {
		t.Errorf("RoleList() error = %v", err)
	}
}

// TestRoleAdd 测试添加角色
func TestRoleAdd(t *testing.T) {
	if err := RoleAdd("test"); err != nil {
		t.Errorf("RoleAdd() error = %v", err)
	}
}

// TestRoleGet 测试查看角色信息
func TestRoleGet(t *testing.T) {
	if err := RoleGet("test"); err != nil {
		t.Errorf("RoleGet() error = %v", err)
	}
}

// TestRoleDelete 测试删除角色
func TestRoleDelete(t *testing.T) {
	if err := RoleDelete("test"); err != nil {
		t.Errorf("RoleDelete() error = %v", err)
	}
}

// TestRoleGrantPermission 测试赋予角色权限
func TestRoleGrantPermission(t *testing.T) {
	role := "test"
	// 为角色 test 赋予 a 的只读权限
	if err := RoleGrantPermission(role, "a", "", clientv3.PermissionType(clientv3.PermRead)); err != nil {
		t.Errorf("RoleGrantPermission() error = %v", err)
	}
	// 为角色 test 赋予 b 的只写权限
	if err := RoleGrantPermission(role, "b", "", clientv3.PermissionType(clientv3.PermWrite)); err != nil {
		t.Errorf("RoleGrantPermission() error = %v", err)
	}
	// 为角色 test 赋予 c 的读写权限
	if err := RoleGrantPermission(role, "c", "", clientv3.PermissionType(clientv3.PermReadWrite)); err != nil {
		t.Errorf("RoleGrantPermission() error = %v", err)
	}
	_ = RoleGet(role)
}

// TestRoleRevokePermission 测试撤销角色权限
func TestRoleRevokePermission(t *testing.T) {
	role := "test"
	// 撤销赋予角色 test 赋予 a 的只读权限
	if err := RoleRevokePermission(role, "a", ""); err != nil {
		t.Errorf("RoleRevokePermission() error = %v", err)
	}
	// 撤销赋予角色 test 赋予 b 的只写权限
	if err := RoleRevokePermission(role, "b", ""); err != nil {
		t.Errorf("RoleRevokePermission() error = %v", err)
	}
	_ = RoleGet(role)
}
