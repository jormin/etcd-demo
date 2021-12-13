package base

import (
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// TestUserAdd 测试添加用户
func TestUserAdd(t *testing.T) {
	if err := UserAdd("test", "111111"); err != nil {
		t.Errorf("UserAdd() error = %v", err)
	}
}

// TestUserAddWithOptions 测试使用选项添加用户
func TestUserAddWithOptions(t *testing.T) {
	options := &clientv3.UserAddOptions{}
	options.NoPassword = true
	if err := UserAddWithOptions("test2", "", options); err != nil {
		t.Errorf("UserAdd() error = %v", err)
	}
}

// TestUserChangePasswd 测试修改密码
func TestUserChangePasswd(t *testing.T) {
	if err := UserChangePasswd("test", "222222"); err != nil {
		t.Errorf("UserAdd() error = %v", err)
	}
}

// TestUserGet 测试查询用户信息
func TestUserGet(t *testing.T) {
	if err := UserGet("test"); err != nil {
		t.Errorf("UserGet() error = %v", err)
	}
}

// TestUserGrantRole 测试赋予用户权限
func TestUserGrantRole(t *testing.T) {
	if err := UserGrantRole("test", "root"); err != nil {
		t.Errorf("UserGrantRole() error = %v", err)
	}
}

// TestUserRevokeRole 测试撤销用户角色
func TestUserRevokeRole(t *testing.T) {
	if err := UserRevokeRole("test", "root"); err != nil {
		t.Errorf("UserRevokeRole() error = %v", err)
	}
}

// TestUserList 测试获取用户列表
func TestUserList(t *testing.T) {
	if err := UserList(); err != nil {
		t.Errorf("UserList() error = %v", err)
	}
}

// TestUserDelete 测试删除用户
func TestUserDelete(t *testing.T) {
	if err := UserDelete("test"); err != nil {
		t.Errorf("UserDelete() error = %v", err)
	}
}
