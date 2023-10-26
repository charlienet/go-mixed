package idgenerator

import "github.com/charlienet/go-mixed/idGenerator/store"

// 序列存储分配器
type storage interface {
	MachineCode() int64                                  // 当前机器码
	UpdateMachineCode(max int64) (int64, error)          // 更新机器标识
	Assign(min, max, step int64) (*store.Segment, error) // 分配号段
	Close()
}
