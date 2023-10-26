package idgenerator

import (
	"time"

	"github.com/charlienet/go-mixed/idGenerator/store"
	"github.com/charlienet/go-mixed/mathx"
	"github.com/charlienet/go-mixed/redis"
)

const (
	defaultDoubleBufferStep int64 = 50
)

var DefaultStartTimeStamp = time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local).Unix()

type opt func(*idGenerator) error

type Option struct {
	TimeFormat    string
	SerialLength  int
	MachineLength int
}

type idGenerator struct {
	store    storage       // 外部存储
	formater formater      // 格式化器
	buffer   *doubleBuffer // 序列缓冲
}

func WithMem(machineCode int64) opt {
	return WithStore(store.NewMemStore(machineCode))
}

func WithRedis(key string, rdb redis.Client) opt {
	return WithStore(store.NewRedisStore(key, rdb))
}

func WithStore(s storage) opt {
	return func(ig *idGenerator) error {
		ig.store = s
		return nil
	}
}

func WithDecimalFormater(format string, serialLength, machineLength int) opt {
	return func(ig *idGenerator) error {
		f, err := newDecimalFormater(format, serialLength, machineLength)
		if err != nil {
			return err
		}

		ig.formater = f
		return nil
	}
}

func WithBinaryFormatter(start int64, serialLength, machineLength int64) opt {
	return func(ig *idGenerator) error {
		f, err := newBinaryFormatter(start, serialLength, machineLength)
		if err != nil {
			return err
		}

		ig.formater = f
		return nil
	}
}

func New(opts ...opt) (*idGenerator, error) {
	g := &idGenerator{}

	for _, o := range opts {
		err := o(g)
		if err != nil {
			return nil, err
		}
	}

	if g.store == nil {
		g.store = store.NewMemStore(0)
	}

	_, err := g.store.UpdateMachineCode(g.formater.MaxinalMachineCode()) // 初始化机器码
	if err != nil {
		return nil, err
	}

	g.buffer = newDoubleBuffer(g.obtain) // 初始化序列缓冲

	return g, nil
}

func (g *idGenerator) WithRedis(key string, rdb redis.Client) *idGenerator {
	return g.WithStore(store.NewRedisStore(key, rdb))
}

func (g *idGenerator) WithStore(s storage) *idGenerator {
	g.store = s
	return g
}

func (g *idGenerator) Next() Id {
	serial, reback := g.buffer.allot()
	id := g.formater.Format(g.store.MachineCode(), serial, reback)

	return id
}

func (g *idGenerator) Close() {
	if g.store != nil {
		g.store.Close()
	}
}

func (g *idGenerator) obtain() (*store.Segment, error) {
	step := mathx.Min(defaultDoubleBufferStep, g.formater.MaximalSequence())
	s, err := g.store.Assign(0, g.formater.MaximalSequence(), step)
	if err != nil {
		println("分配失败", err.Error())
	}

	return s, err
}
