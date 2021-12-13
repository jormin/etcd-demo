package base

import "testing"

// TestAuthStatus 测试获取鉴权开启状态
func TestAuthStatus(t *testing.T) {
	if err := AuthStatus(); err != nil {
		t.Errorf("AuthStatus() error = %v", err)
	}
}

// TestAuthDisable 测试禁用鉴权
func TestAuthDisable(t *testing.T) {
	if err := AuthDisable(); err != nil {
		t.Errorf("AuthDisable() error = %v", err)
	}
}

// TestAuthEnable 测试开启鉴权
func TestAuthEnable(t *testing.T) {
	if err := AuthEnable(); err != nil {
		t.Errorf("AuthEnable() error = %v", err)
	}
}
