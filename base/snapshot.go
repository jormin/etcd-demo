package base

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/snapshot"
)

// SnapshotSave 生成快照
func SnapshotSave(cfg clientv3.Config, dbPath string) error {
	err := snapshot.Save(ctx, nil, cfg, dbPath)
	if err != nil {
		return err
	}
	return nil
}
