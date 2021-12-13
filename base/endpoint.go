package base

import (
	"github.com/jormin/etcd-demo/pkg/helper"
)

// Status 查看状态
func Status() error {
	for _, v := range cli.Endpoints() {
		res, err := cli.Status(ctx, v)
		if err != nil {
			return err
		}
		helper.PrintJSON(res)
	}
	return nil
}

// HashKV 查看节点历史hash
func HashKV() error {
	for _, v := range cli.Endpoints() {
		res, err := cli.HashKV(ctx, v, 0)
		if err != nil {
			return err
		}
		helper.PrintJSON(res)
	}
	return nil
}
