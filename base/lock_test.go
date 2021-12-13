package base

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestMutex 测试分布式锁
func TestMutex(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		time.Sleep(time.Second)
		go func(i int) {
			defer wg.Done()
			mutex, err := Mutex(5, "/lock/foo")
			if err != nil {
				t.Errorf("Mutex() error = %v", err)
			}
			tmpCtx, cancelFun := context.WithTimeout(context.Background(), time.Second*5)
			defer cancelFun()
			err = mutex.Lock(tmpCtx)
			if err != nil {
				t.Errorf("Lock() error = %v", err)
			}
			fmt.Printf("%d got lock: %+v\n", i, mutex)
			err = mutex.Unlock(ctx)
			if err != nil {
				t.Errorf("Unlock() error = %v", err)
			}
			fmt.Printf("%d unlock\n", i)
		}(i)
	}
	wg.Wait()
}
