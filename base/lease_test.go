package base

import (
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// TestLeaseGrant 测试生成租约
func TestLeaseGrant(t *testing.T) {
	if err := LeaseGrant(300); err != nil {
		t.Errorf("LeaseGrant() error = %v", err)
	}
}

// TestLeaseTimeToLive 测试获取租约信息
func TestLeaseTimeToLive(t *testing.T) {
	if err := LeaseTimeToLive(2228857047456997321, clientv3.WithAttachedKeys()); err != nil {
		t.Errorf("LeaseTimeToLive() error = %v", err)
	}
}

// TestLeaseKeepAlive 测试保持租约不过期
func TestLeaseKeepAlive(t *testing.T) {
	if err := LeaseKeepAlive(2228857047456997321); err != nil {
		t.Errorf("LeaseKeepAlive() error = %v", err)
	}
}

// TestLeaseKeepAliveOnce 测试更新租约一次
func TestLeaseKeepAliveOnce(t *testing.T) {
	if err := LeaseKeepAliveOnce(2228857047456997321); err != nil {
		t.Errorf("LeaseKeepAliveOnce() error = %v", err)
	}
}

// TestLeaseList 测试获取租约列表
func TestLeaseList(t *testing.T) {
	if err := LeaseList(); err != nil {
		t.Errorf("LeaseList() error = %v", err)
	}
}

// TestLeaseRevoke 测试撤销租约
func TestLeaseRevoke(t *testing.T) {
	if err := LeaseRevoke(2228857047456997321); err != nil {
		t.Errorf("LeaseRevoke() error = %v", err)
	}
}
