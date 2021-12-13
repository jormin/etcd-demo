package base

import (
	"testing"

	"github.com/jormin/etcd-demo/pkg/etcd"
)

// TestSnapshotSave 测试生成快照
func TestSnapshotSave(t *testing.T) {
	cfg := etcd.GetConfig()
	cfg.Endpoints = []string{"127.0.0.1:12379"}
	if err := SnapshotSave(cfg, "snapshot.db"); err != nil {
		t.Errorf("SnapshotSave() error = %v", err)
	}
}
