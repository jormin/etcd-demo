package base

import (
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// TestPut 测试设置
func TestPut(t *testing.T) {
	if err := Put("foo", "bar"); err != nil {
		t.Errorf("Put() error = %v", err)
	}
	res, _ := cli.Grant(ctx, 10)
	if err := Put("a", "1", clientv3.WithLease(res.ID)); err != nil {
		t.Errorf("Put() error = %v", err)
	}
	if err := Put("a", "2", clientv3.WithPrevKV()); err != nil {
		t.Errorf("Put() error = %v", err)
	}
}

// TestGet 测试获取
func TestGet(t *testing.T) {
	if err := Get("a"); err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if err := Get("a", clientv3.WithCountOnly()); err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if err := Get("foo", clientv3.WithPrefix()); err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if err := Get(
		"foo", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByValue, clientv3.SortDescend),
	); err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if err := Get("foo", clientv3.WithPrefix(), clientv3.WithLimit(1)); err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if err := Get("foo", clientv3.WithPrefix(), clientv3.WithKeysOnly()); err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if err := Get("foo", clientv3.WithFromKey()); err != nil {
		t.Errorf("Get() error = %v", err)
	}
}

// TestDelete 测试删除
func TestDelete(t *testing.T) {
	if err := Delete("foo", clientv3.WithFromKey()); err != nil {
		t.Errorf("Delete() error = %v", err)
	}
}

// TestWatch 测试监听
func TestWatch(t *testing.T) {
	Watch("foo", clientv3.WithPrefix(), clientv3.WithPrevKV())
}
