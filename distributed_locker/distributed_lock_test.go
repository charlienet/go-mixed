package locker

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/tests"
)

func TestDistributedLock(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {
		lock := NewDistributedLocker(context.Background(), "lock_test", rdb)
		lock.Lock()
		lock.Unlock()
	})
}

func TestConcurrence(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {

		count := 5
		var wg sync.WaitGroup
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func(i int) {
				defer wg.Done()

				locker := NewDistributedLocker(context.Background(), "lock_test", rdb)
				for n := 0; n < 5; n++ {
					locker.Lock()
					t.Logf("协程%d获取到锁", i)
					time.Sleep(time.Second)

					t.Logf("协程%d释放锁", i)
					locker.Unlock()
				}
			}(i)
		}

		wg.Wait()
		log.Println("所有任务完成")
	})

}

func TestTwoLocker(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {
		l1 := NewDistributedLocker(context.Background(), "lock_test", rdb)
		l2 := NewDistributedLocker(context.Background(), "lock_test", rdb)

		go func() {
			l1.Lock()
			println("l1 获取锁")
		}()

		go func() {
			l2.Lock()
			println("l2 获取锁")
		}()

		time.Sleep(time.Second * 20)

		l1.Unlock()
		l2.Unlock()
	})
}

func TestDistributediTryLock(t *testing.T) {

	tests.RunOnRedis(t, func(client redis.Client) {
		lock := NewDistributedLocker(context.Background(), "lock_test", client)
		l := lock.TryLock()
		t.Log("尝试加锁结果:", l)

		time.Sleep(time.Second * 20)
		lock.Unlock()
	})

}

func TestLocker(t *testing.T) {
}
