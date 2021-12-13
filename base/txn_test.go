package base

import (
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// TestTxn 测试事务
func TestTxn(t *testing.T) {
	_ = Put("a", "1")
	cmp := []clientv3.Cmp{clientv3.Compare(clientv3.Value("a"), "=", "1")}
	success := []clientv3.Op{clientv3.OpPut("res", "true")}
	fail := []clientv3.Op{clientv3.OpPut("res", "false")}
	err := Txn(cmp, success, fail)
	if err != nil {
		t.Errorf("Txn() error = %v", err)
	}
	_ = Get("res")
}
