package base

import (
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// TestAlarmList 测试获取警报列表
func TestAlarmList(t *testing.T) {
	if err := AlarmList(); err != nil {
		t.Errorf("AlarmList() error = %v", err)
	}
}

// TestAlarmDisarm 测试解除警报
func TestAlarmDisarm(t *testing.T) {
	if err := AlarmDisarm(&clientv3.AlarmMember{}); err != nil {
		t.Errorf("AlarmDisarm() error = %v", err)
	}
}
