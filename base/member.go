package base

import "github.com/jormin/etcd-demo/pkg/helper"

// MemberList 节点列表
func MemberList() error {
	res, err := cli.MemberList(ctx)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// MemberAdd 添加常规节点
func MemberAdd(peerAddrs []string) error {
	res, err := cli.MemberAdd(ctx, peerAddrs)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// MemberAddLearner 添加学习者节点
func MemberAddLearner(peerAddrs []string) error {
	res, err := cli.MemberAddAsLearner(ctx, peerAddrs)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// MemberUpdate 成员更新
func MemberUpdate(id uint64, peerAddrs []string) error {
	res, err := cli.MemberUpdate(ctx, id, peerAddrs)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// MemberPromote 提升学习者节点为常规节点
func MemberPromote(id uint64) error {
	res, err := cli.MemberPromote(ctx, id)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// MemberRemove 删除成员
func MemberRemove(id uint64) error {
	res, err := cli.MemberRemove(ctx, id)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}

// MoveLeader 提升为leader
func MoveLeader(id uint64) error {
	res, err := cli.MoveLeader(ctx, id)
	if err != nil {
		return err
	}
	helper.PrintJSON(res)
	return nil
}
