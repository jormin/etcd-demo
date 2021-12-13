package base

import "testing"

// TestHashKV 测试获取节点历史hash
func TestHashKV(t *testing.T) {
	if err := HashKV(); err != nil {
		t.Errorf("HashKV() error = %v", err)
	}
}

// TestStatus 测试获取状态
func TestStatus(t *testing.T) {
	if err := Status(); err != nil {
		t.Errorf("HashKV() error = %v", err)
	}
}
