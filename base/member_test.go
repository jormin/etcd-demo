package base

import "testing"

// TestMemberList 测试获取成员列表
func TestMemberList(t *testing.T) {
	if err := MemberList(); err != nil {
		t.Errorf("MemberList() error = %v", err)
	}
}

// TestMemberAdd 测试添加常规成员
func TestMemberAdd(t *testing.T) {
	// 添加后需要手动执行下述命令创建 etcd-04
	// etcd --data-dir=/Users/Jormin/etcd/data/etcd4 --name etcd-04 \
	// --initial-advertise-peer-urls http://127.0.0.1:42380 --listen-peer-urls http://127.0.0.1:42380 \
	// --advertise-client-urls http://127.0.0.1:42379 --listen-client-urls http://127.0.0.1:42379 \
	// --initial-cluster etcd-01=http://127.0.0.1:12380,etcd-04=http://127.0.0.1:42380,etcd-02=http://127.0.0.1:22380,etcd-03=http://127.0.0.1:32380 \
	// --initial-cluster-state existing \
	// --initial-cluster-token etcd-cluster
	if err := MemberAdd([]string{"http://127.0.0.1:42380"}); err != nil {
		t.Errorf("MemberAdd() error = %v", err)
	}
}

// TestMemberAddLearner 测试添加学习者成员
func TestMemberAddLearner(t *testing.T) {
	// 添加后需要手动执行下述命令创建 etcd-05
	// etcd --data-dir=/Users/Jormin/etcd/data/etcd5 --name etcd-05 \
	// --initial-advertise-peer-urls http://127.0.0.1:52380 --listen-peer-urls http://127.0.0.1:52380 \
	// --advertise-client-urls http://127.0.0.1:52379 --listen-client-urls http://127.0.0.1:52379 \
	// --initial-cluster etcd-01=http://127.0.0.1:12380,etcd-05=http://127.0.0.1:52380,etcd-04=http://127.0.0.1:42380,etcd-02=http://127.0.0.1:22380,etcd-03=http://127.0.0.1:32380 \
	// --initial-cluster-state existing \
	// --initial-cluster-token etcd-cluster
	if err := MemberAddLearner([]string{"http://127.0.0.1:52380"}); err != nil {
		t.Errorf("MemberAddLearner() error = %v", err)
	}
}

// TestMemberUpdate 测试更新成员
func TestMemberUpdate(t *testing.T) {
	if err := MemberUpdate(11266755900766947019, []string{"http://127.0.0.1:52380"}); err != nil {
		t.Errorf("MemberUpdate() error = %v", err)
	}
}

// TestMemberPromote 测试提升学习者为常规节点
func TestMemberPromote(t *testing.T) {
	if err := MemberPromote(11749075745249145082); err != nil {
		t.Errorf("MemberPromote() error = %v", err)
	}
}

// TestMemberRemove 测试删除成员
func TestMemberRemove(t *testing.T) {
	if err := MemberRemove(3988597216970387585); err != nil {
		t.Errorf("MemberRemove() error = %v", err)
	}
}

// TestMoveLeader 测试提升为leader
func TestMoveLeader(t *testing.T) {
	if err := MoveLeader(15985099378392235387); err != nil {
		t.Errorf("MoveLeader() error = %v", err)
	}
}
