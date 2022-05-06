package pool_test

import (
	"testing"

	"github.com/charlienet/go-mixed/pool"
)

type PoolObject struct {
	Name string
}

func TestPool(t *testing.T) {
	p := pool.NewPool[PoolObject](10)
	o := p.Get()
	o.Name = "abc"
	t.Logf("%p", &o)

	p.Put(o)
	o2 := p.Get()
	t.Logf("取出对象:%s %p", o2, &o2)
}

func TestPoolSize(t *testing.T) {
	p := pool.NewPool[PoolObject](10)
	for i := 0; i < 15; i++ {
		o := p.Get()
		t.Logf("%02d 取出对象:%p  %v  %s", i, &o, o, o.Name)

		if i%2 == 0 {
			p.Put(o)
		}
	}
}

func TestPut(t *testing.T) {
	p := pool.NewPool[PoolObject](10)
	for i := 0; i < 15; i++ {
		p.Put(PoolObject{})
	}

	t.Logf("%+v", *p)
}

func BenchmarkPool(b *testing.B) {
	p := pool.NewPool[int](100)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})
}

func BenchmarkPoolNew(b *testing.B) {
	p := pool.NewPoolWithNew(100, func() int {
		return 100
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(p.Get())
		}
	})
}
