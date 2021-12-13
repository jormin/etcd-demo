package base

import (
	"github.com/jormin/etcd-demo/pkg/helper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Txn 事务
func Txn(cmp []clientv3.Cmp, success []clientv3.Op, fail []clientv3.Op) error {
	txn := cli.Txn(ctx)
	for _,cs := range cmp {
		txn.If(cs)
	}
	for _,op := range success {
		txn.Then(op)
	}
	for _,op := range fail {
		txn.Else(op)
	}
	res, err := txn.Commit()
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}
