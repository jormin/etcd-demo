package base

import (
	"go.etcd.io/etcd/client/v3/concurrency"
)

// Mutex 分布式锁
func Mutex(ttl int64, pfx string) (*concurrency.Mutex, error) {
	res, err := cli.Grant(ctx, ttl)
	if err != nil {
		return nil, err
	}
	s, err := concurrency.NewSession(
		cli, concurrency.WithContext(ctx), concurrency.WithTTL(int(ttl)), concurrency.WithLease(res.ID),
	)
	if err != nil {
		return nil, err
	}
	return concurrency.NewMutex(s, pfx), nil
}
