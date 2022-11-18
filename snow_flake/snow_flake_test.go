package snowflake

import (
	"testing"

	"github.com/charlienet/go-mixed/sets"
)

func TestGet(t *testing.T) {
	s := CreateSnowflake(2)
	t.Log(s.GetId())
}

func TestGetId(t *testing.T) {
	s := CreateSnowflake(22)
	for i := 0; i < 100; i++ {
		t.Log(s.GetId())
	}
}

func TestMutiGetId(t *testing.T) {
	s := CreateSnowflake(11)
	for i := 0; i < 100000; i++ {
		s.GetId()
	}
}

func TestMutiConflict(t *testing.T) {
	set := sets.NewHashSet[int64]()
	s := CreateSnowflake(11)
	for i := 0; i < 10000000; i++ {
		id := s.GetId()
		if set.Contains(id) {
			t.Fatal("失败，生成重复数据")
		}

		set.Add(id)
	}
}

func BenchmarkGetId(b *testing.B) {
	s := CreateSnowflake(11)
	for i := 0; i < b.N; i++ {
		s.GetId()
	}
}

func BenchmarkMutiGetId(b *testing.B) {
	s := CreateSnowflake(11)
	set := sets.NewHashSet[int64]().Sync()
	b.RunParallel(func(p *testing.PB) {
		for i := 0; p.Next(); i++ {
			id := s.GetId()

			if set.Contains(id) {
				b.Fatal("标识重复", id)
			}

			set.Add(id)
		}
	})
}
